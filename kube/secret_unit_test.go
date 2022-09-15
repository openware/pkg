package kube

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSecret_CreateSecret(t *testing.T) {
	t.Run("error for already exists secret", func(t *testing.T) {
		mockSecret := MockSecret("mock-secret", mockNamespace, map[string][]byte{
			"secret": []byte("supersecret"),
		})
		client := NewMockClient(mockSecret)
		err := client.CreateSecret("mock-secret", mockNamespace, "", make(map[string]interface{}))

		assert.Error(t, err)
		assert.True(t, errors.IsAlreadyExists(err))
	})

	t.Run("should create secret", func(t *testing.T) {
		client := NewMockClient()
		err := client.CreateSecret("mock-secret", mockNamespace, "", make(map[string]interface{}))

		assert.NoError(t, err)
	})
}

func TestSecret_CreateEmptySecret(t *testing.T) {
	t.Run("no error for secret with empty field", func(t *testing.T) {
		mockSecret := MockSecret("mock-secret", mockNamespace, map[string][]byte{
			"secret": []byte("supersecret"),
		})
		client := NewMockClient(mockSecret)
		err := client.CreateSecret("mock-secret-test", mockNamespace, "", make(map[string]interface{}))

		assert.NoError(t, err)
	})
}

func TestSecret_UpdateSecret(t *testing.T) {
	t.Run("should create secret if not exists", func(t *testing.T) {
		client := NewMockClient()
		err := client.UpdateSecret("mock-not-found-secret", mockNamespace, make(map[string]interface{}), KeepSecret)

		assert.NoError(t, err)

		cs := client.Client.CoreV1().Secrets(mockNamespace)
		result, err := cs.Get(context.TODO(), "mock-not-found-secret", metav1.GetOptions{})

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("should update existing secret", func(t *testing.T) {
		mockSecret := MockSecret("mock-secret", mockNamespace, map[string][]byte{
			"secret":  []byte("supersecret"),
			"version": []byte("1"),
		})
		expectSecret := make(map[string]interface{})
		expectSecret["secret"] = "updated"
		client := NewMockClient(mockSecret)
		err := client.UpdateSecret("mock-secret", mockNamespace, expectSecret, KeepSecret)

		assert.NoError(t, err)

		cs := client.Client.CoreV1().Secrets(mockNamespace)
		result, err := cs.Get(context.TODO(), "mock-secret", metav1.GetOptions{})
		assert.NoError(t, err)
		actual := string(result.Data["secret"])
		version := string(result.Data["version"])
		assert.Equal(t, "updated", actual)
		assert.Equal(t, "1", version)
	})

	t.Run("should replace secret", func(t *testing.T) {
		mockSecret := MockSecret("mock-secret", mockNamespace, map[string][]byte{
			"secret": []byte("supersecret"),
			"extra":  []byte("data"),
		})
		expectSecret := make(map[string]interface{})
		expectSecret["secret"] = "updated"
		client := NewMockClient(mockSecret)
		err := client.UpdateSecret("mock-secret", mockNamespace, expectSecret, ReplaceSecret)

		assert.NoError(t, err)

		cs := client.Client.CoreV1().Secrets(mockNamespace)
		result, err := cs.Get(context.TODO(), "mock-secret", metav1.GetOptions{})
		assert.NoError(t, err)
		actual := string(result.Data["secret"])
		assert.Equal(t, "updated", actual)

		// extra secret should not be found
		_, ok := result.Data["extra"]
		assert.False(t, ok)
	})
}

func TestSecret_ReadSecret(t *testing.T) {
	t.Run("error reading from not exists secret", func(t *testing.T) {
		client := NewMockClient()
		_, err := client.ReadSecret("mock-secret", mockNamespace)

		assert.Error(t, err)
		assert.True(t, errors.IsNotFound(err))
	})

	t.Run("should read existing secret", func(t *testing.T) {
		mockSecret := MockSecret("mock-secret", mockNamespace, map[string][]byte{
			"secret": []byte("supersecret"),
		})
		client := NewMockClient(mockSecret)
		result, err := client.ReadSecret("mock-secret", mockNamespace)

		assert.NoError(t, err)
		actual := string(result["secret"])
		assert.Equal(t, "supersecret", actual)
	})
}

func TestPullSecret_GetAnnotatedPullSecrets(t *testing.T) {
	managed := map[string]string{
		"app.kubernetes.io/managed-by": "opendax",
	}

	mockSecret1 := MockSecret("mock-secret-1", mockNamespace, map[string][]byte{
		"secret": []byte("supersecret"),
	})
	mockSecret1.Type = corev1.SecretTypeDockerConfigJson
	mockSecret1.Annotations = managed
	mockSecret2 := MockSecret("mock-secret-2", mockNamespace, map[string][]byte{
		"secret": []byte("supersecret"),
	})
	mockSecret2.Type = corev1.SecretTypeDockerConfigJson
	mockSecret2.Annotations = managed

	t.Run("should get pull secrets", func(t *testing.T) {
		client := NewMockClient(mockSecret1, mockSecret2)

		secrets, err := client.GetAnnotatedPullSecrets(mockNamespace, managed)
		assert.NoError(t, err)
		assert.NotEmpty(t, secrets)

		assert.Equal(t, "mock-secret-1", secrets[0].Name)
		assert.Equal(t, "mock-secret-2", secrets[1].Name)
	})
}
