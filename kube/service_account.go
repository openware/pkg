package kube

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) CreateServiceAccount(name, namespace string) (*corev1.ServiceAccount, error) {
	saClient := c.Client.CoreV1().ServiceAccounts(namespace)

	sa := &corev1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Annotations: map[string]string{
				"helm.sh/resource-policy": "keep",
			},
		},
	}

	sa, err := saClient.Create(context.TODO(), sa, metav1.CreateOptions{})
	if err != nil {
		return nil, err
	}

	return sa, nil
}

func (c *K8sClient) GetServiceAccount(name, namespace string) (*corev1.ServiceAccount, error) {
	saClient := c.Client.CoreV1().ServiceAccounts(namespace)

	sa, err := saClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return sa, nil
}
