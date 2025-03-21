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

package flags

import "fmt"

const (
	viperInfraPrefix             = "infra"
	viperKubernetesClusterPrefix = "k8s"
	viperS3Prefix                = "s3"
	viperBuildPrefix             = "build"
	viperScanPrefix              = "scan"
	viperSignPrefix              = "sign"
)

var (
	viperOpenStackPrefix = fmt.Sprintf("%s.openstack", viperInfraPrefix)
	viperKubeVirtPrefix  = fmt.Sprintf("%s.kubevirt", viperInfraPrefix)
	viperGpuPrefix       = fmt.Sprintf("%s.gpu", viperBuildPrefix)
	viperVaultPrefix     = fmt.Sprintf("%s.vault", viperSignPrefix)
	viperGeneratePrefix  = fmt.Sprintf("%s.generate", viperSignPrefix)
	viperSinglePrefix    = fmt.Sprintf("%s.single", viperScanPrefix)
	viperMultiplePrefix  = fmt.Sprintf("%s.multiple", viperScanPrefix)
)
