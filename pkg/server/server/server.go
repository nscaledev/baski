/*
Copyright 2024 Drewbernetes.

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

package server

import (
	"fmt"
	"github.com/drewbernetes/baski/pkg/server/generated"
	"github.com/drewbernetes/baski/pkg/server/handler"
	simple_s3 "github.com/drewbernetes/simple-s3"
	"github.com/gorilla/mux"
	"net/http"
)

type Options struct {
	ListenAddress string
	ListenPort    int32
	Endpoint      string
	AccessKey     string
	SecretKey     string
	Region        string
	Bucket        string
	EnableDogKat  bool
	DogKatBucket  string
	CloudName     string
}

type Server struct {
	Options Options
}

func (s *Server) NewServer(dev bool) (*http.Server, error) {

	middleware := []generated.MiddlewareFunc{}
	r := mux.NewRouter()

	r.Use(mux.CORSMethodMiddleware(r))
	if dev {
		middleware = append(middleware, corsAllowOriginAllMiddleware)
	}

	baskiS3, err := simple_s3.New(s.Options.Endpoint, s.Options.AccessKey, s.Options.SecretKey, s.Options.Bucket, s.Options.Region)

	if err != nil {
		return nil, err
	}

	var dogKatS3 *simple_s3.S3
	if s.Options.EnableDogKat {
		dogKatS3, err = simple_s3.New(s.Options.Endpoint, s.Options.AccessKey, s.Options.SecretKey, s.Options.DogKatBucket, s.Options.Region)
		if err != nil {
			return nil, err
		}

	}
	handlers := handler.New(baskiS3, dogKatS3, s.Options.CloudName)

	options := generated.GorillaServerOptions{
		BaseRouter:  r,
		Middlewares: middleware,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Options.ListenAddress, s.Options.ListenPort),
		Handler: generated.HandlerWithOptions(handlers, options),
	}

	return server, nil
}

// corsAllowOriginAllMiddleware sets the header for Access-Control-Allow-Origin = "*"
func corsAllowOriginAllMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		next(w, r)
	}
}
