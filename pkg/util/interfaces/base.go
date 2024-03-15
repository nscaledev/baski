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
