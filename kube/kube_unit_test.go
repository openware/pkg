package kube

import (
	"fmt"
	"net"
	"os"
	"testing"

	b64 "encoding/base64"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/cert"
)

func TestKube_NewClient(t *testing.T) {
	t.Run("error for malformed config", func(t *testing.T) {
		fs, err := os.CreateTemp("", "tmp-config-")
		assert.NoError(t, err)

		_, err = fs.WriteString(`apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K
    server: https://localhost:16443
  name: mock-cluster
contexts:
- context:
    cluster: mock-cluster
    user: admin
  name: mock
current-context: mock
kind: Config
preferences: {}
users:
- name: admin
  user:
    token: mock-token`)
		assert.NoError(t, err)

		defer fs.Close()
		defer os.Remove(fs.Name())

		cfg, err := clientcmd.BuildConfigFromFlags("", fs.Name())
		assert.NoError(t, err)

		_, err = NewClient(cfg)
		assert.Error(t, err)
	})

	t.Run("should create client", func(t *testing.T) {
		fs, err := os.CreateTemp("", "tmp-config-")
		assert.NoError(t, err)

		crt, _, err := cert.GenerateSelfSignedCertKey("localhost", []net.IP{}, []string{})
		assert.NoError(t, err)

		mockCert := b64.URLEncoding.EncodeToString(crt)
		conf := fmt.Sprintf(`apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: %s
    server: https://localhost:16443
  name: mock-cluster
contexts:
- context:
    cluster: mock-cluster
    user: admin
  name: mock
current-context: mock
kind: Config
preferences: {}
users:
- name: admin
  user:
    token: mock-token`, mockCert)

		_, err = fs.WriteString(conf)
		assert.NoError(t, err)

		defer fs.Close()
		defer os.Remove(fs.Name())

		cfg, err := clientcmd.BuildConfigFromFlags("", fs.Name())
		assert.NoError(t, err)

		_, err = NewClient(cfg)
		assert.NoError(t, err)
	})
}

func TestKube_CreateNamespace(t *testing.T) {
	t.Run("should skip for existing namespace", func(t *testing.T) {
		mockNS := &v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: mockNamespace,
			},
		}
		client := NewMockClient(mockNS)
		err := client.CreateNamespace(mockNamespace)

		assert.NoError(t, err)
	})

	t.Run("should create namespace", func(t *testing.T) {
		client := NewMockClient()
		err := client.CreateNamespace(mockNamespace)

		assert.NoError(t, err)
	})
}
