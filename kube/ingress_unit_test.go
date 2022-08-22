package kube

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIngress_GetIngress(t *testing.T) {
	t.Run("error for not exists ingress", func(t *testing.T) {
		client := NewMockClient()
		_, err := client.GetIngress(mockNamespace, "mock-ingress")

		assert.Error(t, err)
	})

	t.Run("should get ingress", func(t *testing.T) {
		mockIngress := MockIngress("mock-ingress", mockNamespace)
		client := NewMockClient(mockIngress)
		_, err := client.GetIngress(mockNamespace, "mock-ingress")

		assert.NoError(t, err)
	})
}

func TestIngress_DeleteIngress(t *testing.T) {
	t.Run("error for not exists ingress", func(t *testing.T) {
		client := NewMockClient()
		err := client.DeleteIngress(mockNamespace, "mock-ingress")

		assert.Error(t, err)
	})

	t.Run("should delete ingress", func(t *testing.T) {
		mockIngress := MockIngress("mock-ingress", mockNamespace)
		client := NewMockClient(mockIngress)
		err := client.DeleteIngress(mockNamespace, "mock-ingress")

		assert.NoError(t, err)
	})
}

func TestIngress_ListIngresses(t *testing.T) {
	t.Run("should list ingress", func(t *testing.T) {
		mockIngress := MockIngress("mock-ingress", mockNamespace)
		client := NewMockClient(mockIngress)
		_, err := client.ListIngresses(mockNamespace)

		assert.NoError(t, err)
	})
}
