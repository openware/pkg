package kube

import (
	"context"
	"fmt"
	"os"

	logging "github.com/ipfs/go-log/v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const (
	ToggleAnnotation    = "opendax.io/auto-update"           // enabled/disabled
	ContainerAnnotation = "opendax.io/auto-update-container" // *name*

	CustomIngressAnnotation    = "opendax.io/custom-ingress"    // true/false
	VersionAvailableAnnotation = "opendax.io/version-available" // *version*

	K8sRevisionAnnotation = "deployment.kubernetes.io/revision"
	ManagedByAnnotation   = "app.kubernetes.io/managed-by"
)

var log = logging.Logger("kube")

func getKubeconfigPath() string {
	kubeconfig, ok := os.LookupEnv("KUBECONFIG")
	if ok {
		return kubeconfig
	} else {
		return fmt.Sprintf("%s/.kube/config", homedir.HomeDir())
	}
}

func getImagePullSecret(ipss []corev1.LocalObjectReference) string {
	var ips string
	if len(ipss) > 0 {
		ips = ipss[0].Name
	}

	return ips
}

func flattenLabelSelector(labelSelector *metav1.LabelSelector) (string, error) {
	labelMap, err := metav1.LabelSelectorAsMap(labelSelector)
	if err != nil {
		return "", err
	}

	return labels.SelectorFromSet(labelMap).String(), nil
}

func convertMapInterfaceToString(init map[string]interface{}) map[string][]byte {
	res := make(map[string][]byte)
	for k, v := range init {
		res[k] = []byte(fmt.Sprintf("%v", v))
	}

	return res
}

// Client for all K8s API interactions
type Client struct {
	ConfigPath string
	Config     *rest.Config
	Clientset  kubernetes.Interface
	// Auditor    *audit.Auditor
}

// NewClient returns an initialized Client object
func NewClient() (*Client, error) {
	client := &Client{}
	var err error

	client.Config, err = clientcmd.BuildConfigFromFlags("", getKubeconfigPath())
	if err != nil {
		return nil, err
	}

	client.Clientset, err = kubernetes.NewForConfig(client.Config)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// CreateNamespace creates a namespace with a given name
func (c *Client) CreateNamespace(name string) error {
	nsClient := c.Clientset.CoreV1().Namespaces()

	ns := &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}

	_, err := nsClient.Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			result, err := nsClient.Create(context.TODO(), ns, metav1.CreateOptions{})
			if err != nil {
				return err
			}
			log.Infof("created namespace %s", result.GetObjectMeta().GetName())

		} else {
			return err
		}
	}

	return nil
}
