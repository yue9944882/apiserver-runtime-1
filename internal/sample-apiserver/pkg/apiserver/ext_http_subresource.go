/*
Copyright 2016 The Kubernetes Authors.

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

package apiserver

import (
	"k8s.io/apiserver/pkg/registry/rest"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcerest"
)

func isSubResource(storage rest.Storage) bool {
	_, ok := storage.(resource.SubResource)
	return ok
}

func isHTTPHandlerSubResource(storage rest.Storage) bool {
	if !isSubResource(storage) {
		return false
	}
	if _, ok := storage.(resourcerest.HTTPHandler); !ok {
		return false
	}
	return true
}
