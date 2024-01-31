package interfaces

import (
	"github.com/drewbernetes/baski/pkg/server/generated"
	"net/http"
	"os"
)

//go:generate mockgen -source=base.go -destination=../../mock/base.go -package=mock

type HandlerInterface interface {
	Healthz(w http.ResponseWriter, r *http.Request)
	ApiV1GetScans(w http.ResponseWriter, r *http.Request)
	ApiV1GetScan(w http.ResponseWriter, r *http.Request, imageId generated.ImageID)
	ApiV1GetTests(w http.ResponseWriter, r *http.Request)
	ApiV1GetTest(w http.ResponseWriter, r *http.Request, imageId generated.ImageID)
}

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
