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

package gitRepo

import (
	"github.com/go-git/go-git/v5/plumbing"
	"os"
	"testing"
)

func TestGitClone(t *testing.T) {
	//FIXME: This check is in place until the security branch in this repo go upstream.
	// Until it has been added, we must force users over to this repo as it's the only one that has these new additions.
	repo := "https://github.com/drew-viles/image-builder.git"
	cloneLocation := "/tmp/test"
	err := os.RemoveAll(cloneLocation)
	if err != nil {
		t.Error(err)
		return
	}
	ref := plumbing.ReferenceName("refs/heads/main")
	_, err = GitClone(repo, cloneLocation, ref)
	if err != nil {
		t.Error(err)
		return
	}

	f, err := os.Stat(cloneLocation)
	if err != nil {
		t.Error(err)
		return
	}

	if !f.IsDir() {
		t.Error("expected directory, didn't get a directory")
	}

}
