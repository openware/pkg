package kube

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/util/retry"
)

func (c *K8sClient) ListDeployments(namespace string) (*appsv1.DeploymentList, error) {
	return c.Client.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
}

// UpdateDeployment applies update with given deployment
func (c *K8sClient) UpdateDeployment(deploy *appsv1.Deployment) error {
	deployClient := c.Client.AppsV1().Deployments(deploy.Namespace)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of Deployment before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		if _, err := deployClient.Update(context.TODO(), deploy, metav1.UpdateOptions{}); err != nil {
			return err
		}
		return nil
	})

	if retryErr != nil {
		return fmt.Errorf("deployment update failed: %s", retryErr.Error())
	}

	return nil
}

// ParseDeploymentInfo returns image and its pull secret for a given container in the Deployment
func ParseDeploymentInfo(deploy appsv1.Deployment, conIndx int) (string, string, error) {
	image := deploy.Spec.Template.Spec.Containers[conIndx].Image
	ips := getImagePullSecret(deploy.Spec.Template.Spec.ImagePullSecrets)

	return image, ips, nil
}

func (c *K8sClient) GetDeployment(namespace, name string) (*appsv1.Deployment, error) {
	return c.Client.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func SetDeploymentImage(deploy appsv1.Deployment, image string, conIndx int) {
	deploy.Spec.Template.Spec.Containers[conIndx].Image = image
}

func SetDeploymentImagePullSecret(deploy appsv1.Deployment, imagePullSecret string) {
	imagePullSecrets := []corev1.LocalObjectReference{}
	if imagePullSecret != "" {
		imagePullSecrets = []corev1.LocalObjectReference{{Name: imagePullSecret}}
	}

	deploy.Spec.Template.Spec.ImagePullSecrets = imagePullSecrets
}

func SetDeploymentAnnotation(deploy appsv1.Deployment, annotation, value string) {
	deploy.Annotations[annotation] = value
}

func RemoveDeploymentAnnotation(deploy appsv1.Deployment, annotation, value string) {
	delete(deploy.Annotations, annotation)
}

func GetDeploymentAnnotations(deploy appsv1.Deployment) map[string]string {
	return deploy.Annotations
}

func GetDeploymentResourceVersion(deploy appsv1.Deployment) string {
	return deploy.ResourceVersion
}
