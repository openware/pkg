package kube

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

// CreateSecret creates a K8s secret with a given name from a given map
func (c *K8sClient) CreateSecret(name, namespace string, secType corev1.SecretType, data map[string]interface{}) error {
	secretsClient := c.Client.CoreV1().Secrets(namespace)

	if secType == "" {
		secType = "Opaque"
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Annotations: map[string]string{
				"helm.sh/resource-policy": "keep",
			},
		},
		Data: convertMapInterfaceToString(data),
		Type: secType,
	}

	_, err := secretsClient.Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		return err
	}

	return nil
}

// UpdateOption defines the way to update secrets
type UpdateOption string

const (
	// Keep existing secrets and update with a given secret map
	KeepSecret UpdateOption = "Keep"
	// Replace all secrets by a given secret map
	ReplaceSecret UpdateOption = "Replace"
)

// UpdateSecret updates a K8s secret with a given name from a given map and creates one if it's absent
func (c *K8sClient) UpdateSecret(name, namespace string, data map[string]interface{}, option UpdateOption) error {
	secretsClient := c.Client.CoreV1().Secrets(namespace)

	result, err := secretsClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			return c.CreateSecret(name, namespace, "Opaque", data)
		}

		return err
	}

	byteData := convertMapInterfaceToString(data)
	var resData map[string][]byte
	if option == KeepSecret {
		resData = result.Data
		for k, v := range byteData {
			resData[k] = v
		}
	} else {
		resData = byteData
	}

	result.Data = resData
	result.ObjectMeta.Annotations = map[string]string{
		"helm.sh/resource-policy": "keep",
	}

	_, err = secretsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

// ReadSecret reads a K8s secret with a given name
func (c *K8sClient) ReadSecret(name, namespace string) (map[string][]byte, error) {
	secretsClient := c.Client.CoreV1().Secrets(namespace)

	result, err := secretsClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return result.Data, nil
}

// GetAnnotatedPullSecrets get a K8s secret with a given name and annotations
func (c *K8sClient) GetAnnotatedPullSecrets(namespace string, annotations map[string]string) ([]corev1.Secret, error) {
	secretsClient := c.Client.CoreV1().Secrets(namespace)

	secrets, err := secretsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	ps := []corev1.Secret{}
	for _, secret := range secrets.Items {
		if secret.Type == corev1.SecretTypeDockerConfigJson {
			for key, elem := range annotations {
				if val, ok := secret.Annotations[key]; ok && elem == val {
					ps = append(ps, secret)
				}
			}
		}
	}

	return ps, nil
}

// GetSecrets get a K8s secret with a given name
func (c *K8sClient) GetSecrets(namespace string) ([]corev1.Secret, error) {
	secretsClient := c.Client.CoreV1().Secrets(namespace)

	secrets, err := secretsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	ps := []corev1.Secret{}
	for _, secret := range secrets.Items {
		ps = append(ps, secret)
	}

	return ps, nil
}
