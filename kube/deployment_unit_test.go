package kube

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ContainerAnnotation = "opendax.io/auto-update-container" // *name*
)

func TestDeployment_ListDeployments(t *testing.T) {
	t.Run("return multiple deployments", func(t *testing.T) {
		mockDeployment1 := MockDeployment("mock-deploy-1", mockNamespace)
		mockDeployment2 := MockDeployment("mock-deploy-2", mockNamespace)
		client := NewMockClient(mockDeployment1, mockDeployment2)
		ws, err := client.ListDeployments(mockNamespace)

		assert.NoError(t, err)
		assert.NotEmpty(t, ws)
		assert.Equal(t, "mock-deploy-1", ws.Items[0].Name)
		assert.Equal(t, "mock-deploy-2", ws.Items[1].Name)
	})
}

func TestDeployment_UpdateDeployment(t *testing.T) {
	deploymentName := "mock-deploy"

	t.Run("should update deployment image", func(t *testing.T) {
		containerName := "mock-container"
		mockDeployment := MockDeployment(deploymentName, mockNamespace)
		mockDeployment.Annotations = map[string]string{
			ContainerAnnotation: containerName,
		}
		mockDeployment.Spec = v1.DeploymentSpec{
			Template: MockTemplateSpec(containerName, "finex", ""),
		}
		client := NewMockClient(mockDeployment)

		expected := "frontdex"
		expectedDeployment := MockDeployment(deploymentName, mockNamespace)
		expectedDeployment.Annotations = map[string]string{
			ContainerAnnotation: containerName,
		}
		expectedDeployment.Spec = v1.DeploymentSpec{
			Template: MockTemplateSpec(containerName, expected, ""),
		}
		err := client.UpdateDeployment(expectedDeployment)

		assert.NoError(t, err)

		actual, err := client.Client.AppsV1().Deployments(mockNamespace).Get(
			context.TODO(),
			deploymentName,
			metav1.GetOptions{},
		)

		assert.NoError(t, err)
		assert.Equal(t, expected, actual.Spec.Template.Spec.Containers[0].Image)
	})
}

func TestDeployment_ParseDeployment(t *testing.T) {
	t.Run("should get deployment image and pull secrets", func(t *testing.T) {
		containerName := "mock-container"
		mockDeployment := MockDeployment("mock-deploy", mockNamespace)
		mockDeployment.Annotations = map[string]string{
			ContainerAnnotation: containerName,
		}
		mockDeployment.Spec = v1.DeploymentSpec{
			Template: MockTemplateSpec(containerName, "finex", "mock-pull-secret"),
		}

		image, ips, err := ParseDeploymentInfo(*mockDeployment, 0)
		assert.NoError(t, err)
		assert.Equal(t, "finex", image)
		assert.Equal(t, "mock-pull-secret", ips)
	})
}

func TestDeployment_GetDeployment(t *testing.T) {
	t.Run("error for not exists deployment", func(t *testing.T) {
		client := NewMockClient()
		_, err := client.GetDeployment(mockNamespace, "mock-deploy")

		assert.Error(t, err)
		assert.True(t, errors.IsNotFound(err))
	})

	t.Run("should get deployment", func(t *testing.T) {
		mockDeployment := MockDeployment("mock-deploy", mockNamespace)
		client := NewMockClient(mockDeployment)
		deploy, err := client.GetDeployment(mockNamespace, "mock-deploy")

		assert.NoError(t, err)
		assert.NotNil(t, deploy)
	})
}
