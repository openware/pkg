package kube

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) CreateRoleBinding(name, namespace string, roleRef rbacv1.RoleRef, subject rbacv1.Subject) error {
	roleBindingClient := c.Client.RbacV1().RoleBindings(namespace)

	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Annotations: map[string]string{
				"helm.sh/resource-policy": "keep",
			},
		},
		RoleRef:  roleRef,
		Subjects: []rbacv1.Subject{subject},
	}

	_, err := roleBindingClient.Create(context.TODO(), roleBinding, metav1.CreateOptions{})
	return err
}

func (c *K8sClient) GetRoleBinding(name, namespace string) (*rbacv1.RoleBinding, error) {
	roleBindingClient := c.Client.RbacV1().RoleBindings(namespace)

	roleBinding, err := roleBindingClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return roleBinding, nil
}
