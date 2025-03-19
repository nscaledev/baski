/*
Copyright 2025 Nscale.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	dv_client "kubevirt.io/client-go/generated/containerized-data-importer/clientset/versioned"
	"os"
	"testing"
)

var kubeconfig = `
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: 
    server: https://127.0.0.1:6443
  name: default
contexts:
- context:
    cluster: default
    namespace: default
    user: default
  name: default@default
current-context: default@default
kind: Config
preferences: {}
users:
- name: default
  user:
    client-certificate-data: 
    client-key-data: 
`

func TestNewClient(t *testing.T) {
	kubeconfigPath := "/tmp/kubeconfig"
	err := os.WriteFile(kubeconfigPath, []byte(kubeconfig), 0700)
	if err != nil {
		t.Error(err.Error())
	}

	tc := []struct {
		name   string
		client *KubernetesClient
	}{
		{
			name: "Testing we get a Kubernetes client back",
			client: &KubernetesClient{
				Client:   &kubernetes.Clientset{},
				KubeVirt: &dv_client.Clientset{},
				Config: &rest.Config{
					Host: "https://127.0.0.1:6443",
				},
			},
		},
	}

	for _, v := range tc {
		t.Run(v.name, func(t *testing.T) {
			client, err := NewClient(kubeconfigPath)
			if err != nil {
				t.Error(err.Error())
			}

			//Basic checks for now - much more advanced testing will be required!
			//TODO: Build better tests here than basic comparisons

			if client.Client == nil {
				t.Errorf("exepected a kubernetes ClientSet, got nil\n")
			}
			if client.KubeVirt == nil {
				t.Errorf("exepected a kubevirt ClientSet, got nil\n")
			}
			if client.Config == nil {
				t.Errorf("exepected a rest Config, got nil\n")
			}
		})
	}
}
