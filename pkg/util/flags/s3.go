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

import (
	"fmt"
	"github.com/spf13/viper"
)

type S3Flags struct {
	Endpoint  string
	AccessKey string
	SecretKey string
	Region    string
	IsCeph    bool
}

func (o *S3Flags) SetOptionsFromViper() {
	o.Endpoint = viper.GetString(fmt.Sprintf("%s.endpoint", viperS3Prefix))
	o.AccessKey = viper.GetString(fmt.Sprintf("%s.access-key", viperS3Prefix))
	o.SecretKey = viper.GetString(fmt.Sprintf("%s.secret-key", viperS3Prefix))
	o.Region = viper.GetString(fmt.Sprintf("%s.region", viperS3Prefix))
	o.IsCeph = viper.GetBool(fmt.Sprintf("%s.is-ceph", viperS3Prefix))
}
