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

package interfaces

import (
	"os"
)

//go:generate mockgen -source=base.go -destination=../../mock/base.go -package=mock
type VaultInterface interface {
	Fetch(mountPath, secretPath, data string) ([]byte, error)
}

type SSHInterface interface {
	CopyFromRemoteServer(src, dst string) (*os.File, error)
	SSHClose() error
	SFTPClose() error
}

type S3Interface interface {
	List(string) ([]string, error)
	Fetch(string) ([]byte, error)
	Put(key string, body *os.File) error
}
