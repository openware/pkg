package kube

import (
	"context"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (c *K8sClient) CreateRole(name, namespace string, rules []rbacv1.PolicyRule) error {
	rolesClient := c.Client.RbacV1().Roles(namespace)

	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Annotations: map[string]string{
				"helm.sh/resource-policy": "keep",
			},
		},
		Rules: rules,
	}

	_, err := rolesClient.Create(context.TODO(), role, metav1.CreateOptions{})
	return err
}

func (c *K8sClient) GetRole(name, namespace string) (*rbacv1.Role, error) {
	rolesClient := c.Client.RbacV1().Roles(namespace)

	role, err := rolesClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}

	return role, nil
}

func (c *K8sClient) ListRoles(namespace string) (*rbacv1.RoleList, error) {
	rolesClient := c.Client.RbacV1().Roles(namespace)

	roles, err := rolesClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return roles, nil
}
