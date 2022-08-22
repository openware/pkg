package kube

import (
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestStatefulSet_UpdateStatefulSet(t *testing.T) {
	t.Run("error for malformed label selector", func(t *testing.T) {
		mockSts := MockStatefulSet("mock-sts", mockNamespace)
		mockSts.Spec = v1.StatefulSetSpec{
			Template: MockTemplateSpec("sts-test-container", "mysql", "mock-pull-secret"),
			Selector: &metav1.LabelSelector{
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{
						Operator: "unsupport",
					},
				},
			},
		}
		client := NewMockClient(mockSts)
		err := client.UpdateStatefulSet(mockSts)

		assert.Error(t, err)
	})

	t.Run("should update statefulset image", func(t *testing.T) {
		mockSts := MockStatefulSet("mock-sts", mockNamespace)
		mockSts.Spec = v1.StatefulSetSpec{
			Template: MockTemplateSpec("sts-test-container", "mysql", "mock-pull-secret"),
		}
		client := NewMockClient(mockSts)
		err := client.UpdateStatefulSet(mockSts)

		assert.NoError(t, err)
	})
}

func TestStatefulSet_GetStatefulSet(t *testing.T) {
	t.Run("should get statefulset", func(t *testing.T) {
		mockSts := MockStatefulSet("mock-sts", mockNamespace)
		client := NewMockClient(mockSts)
		_, err := client.GetStatefulSet(mockNamespace, "mock-sts")

		assert.NoError(t, err)
	})
}

func TestStatefulSet_ListStatefulSets(t *testing.T) {
	t.Run("return multiple statefulsets", func(t *testing.T) {
		mockSts1 := MockStatefulSet("mock-sts-1", mockNamespace)
		mockSts2 := MockStatefulSet("mock-sts-2", mockNamespace)
		client := NewMockClient(mockSts1, mockSts2)
		ws, err := client.ListStatefulSets(mockNamespace)

		assert.NoError(t, err)
		assert.NotEmpty(t, ws)
		assert.Equal(t, "mock-sts-1", ws.Items[0].Name)
		assert.Equal(t, "mock-sts-2", ws.Items[1].Name)
	})
}
