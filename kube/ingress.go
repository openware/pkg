package kube

import (
	"context"

	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) GetIngress(namespace, name string) (*netv1.Ingress, error) {
	return c.Client.NetworkingV1().Ingresses(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func (c *K8sClient) DeleteIngress(namespace, name string) error {
	return c.Client.NetworkingV1().Ingresses(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
}

func (c *K8sClient) ListIngresses(namespace string) (*netv1.IngressList, error) {
	return c.Client.NetworkingV1().Ingresses(namespace).List(context.TODO(), metav1.ListOptions{})
}
