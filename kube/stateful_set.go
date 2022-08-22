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

// ListStatefulSets returns a list of stateful sets in a given namespace with enabled auto updates
func (c *K8sClient) ListStatefulSets(namespace string) (*appsv1.StatefulSetList, error) {
	return c.Client.AppsV1().StatefulSets(namespace).List(context.TODO(), metav1.ListOptions{})
}

// UpdateStatefulSet applies update with given deployment
func (c *K8sClient) UpdateStatefulSet(sts *appsv1.StatefulSet) error {
	deployClient := c.Client.AppsV1().StatefulSets(sts.Namespace)
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		// Retrieve the latest version of StatefulSet before attempting update
		// RetryOnConflict uses exponential backoff to avoid exhausting the apiserver
		if _, err := deployClient.Update(context.TODO(), sts, metav1.UpdateOptions{}); err != nil {
			return err
		}

		return nil
	})

	if retryErr != nil {
		return fmt.Errorf("stateful set update failed: %s", retryErr.Error())
	}

	labelSelector, err := FlattenLabelSelector(sts.Spec.Selector)
	if err != nil {
		return err
	}

	if err := c.Client.CoreV1().Pods(sts.Namespace).DeleteCollection(context.TODO(), metav1.DeleteOptions{}, metav1.ListOptions{
		LabelSelector: labelSelector,
	}); err != nil {
		return err
	}

	return nil
}

// ParseStatefulSetInfo returns image and its pull secret for a given container in the StatefulSet
func ParseStatefulSetInfo(sts appsv1.StatefulSet, conIndx int) (string, string, error) {
	image := sts.Spec.Template.Spec.Containers[conIndx].Image
	ips := getImagePullSecret(sts.Spec.Template.Spec.ImagePullSecrets)

	return image, ips, nil
}

func (c *K8sClient) GetStatefulSet(namespace, name string) (*appsv1.StatefulSet, error) {
	return c.Client.AppsV1().StatefulSets(namespace).Get(context.TODO(), name, metav1.GetOptions{})
}

func SetStatefulSetImage(sts appsv1.StatefulSet, image string, conIndx int) {
	sts.Spec.Template.Spec.Containers[conIndx].Image = image
}

func SetStsImagePullSecret(sts appsv1.StatefulSet, imagePullSecret string) {
	imagePullSecrets := []corev1.LocalObjectReference{}
	if imagePullSecret != "" {
		imagePullSecrets = []corev1.LocalObjectReference{{Name: imagePullSecret}}
	}

	sts.Spec.Template.Spec.ImagePullSecrets = imagePullSecrets
}

func SetStsAnnotation(sts appsv1.StatefulSet, annotation, value string) {
	sts.Annotations[annotation] = value
}

func RemoveStsAnnotation(sts appsv1.StatefulSet, annotation, value string) {
	delete(sts.Annotations, annotation)
}

func GetStsAnnotations(sts appsv1.StatefulSet) map[string]string {
	return sts.Annotations
}

func GetStsResourceVersion(sts appsv1.StatefulSet) string {
	return sts.ResourceVersion
}
