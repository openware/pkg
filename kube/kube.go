package kube

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
)

func getImagePullSecret(ipss []corev1.LocalObjectReference) string {
	var ips string
	if len(ipss) > 0 {
		ips = ipss[0].Name
	}

	return ips
}

func FlattenLabelSelector(labelSelector *metav1.LabelSelector) (string, error) {
	labelMap, err := metav1.LabelSelectorAsMap(labelSelector)
	if err != nil {
		return "", err
	}

	return labels.SelectorFromSet(labelMap).String(), nil
}

func convertMapInterfaceToString(init map[string]interface{}) map[string][]byte {
	res := make(map[string][]byte)
	for k, v := range init {
		res[k] = []byte(fmt.Sprintf("%v", v))
	}

	return res
}

// Client for K8s API interactions
type K8sClient struct {
	Client kubernetes.Interface
}

// NewClient returns an initialized K8sClient object
func NewClient(conf *rest.Config) (*K8sClient, error) {
	client := &K8sClient{}
	var err error

	client.Client, err = kubernetes.NewForConfig(conf)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// CreateNamespace creates a namespace with a given name
func (c *K8sClient) CreateNamespace(name string) error {
	nsClient := c.Client.CoreV1().Namespaces()

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	_, err := nsClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			_, err := nsClient.Create(context.TODO(), ns, metav1.CreateOptions{})
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}
