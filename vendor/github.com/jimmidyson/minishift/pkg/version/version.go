/*
Copyright (C) 2016 Red Hat, Inc.

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

package version

import (
	"strings"

	"github.com/blang/semver"
)

// The current version of the minikube and openshift
// This is a private field and should be set when compiling with --ldflags="-X github.com/jimmidyson/minishift/pkg/version.version=vX.Y.Z"
const VersionPrefix = "v"

var version = "v0.0.0-unset"

func GetVersion() string {
	return version
}

func GetSemverVersion() (semver.Version, error) {
	return semver.Make(strings.TrimPrefix(GetVersion(), VersionPrefix))
}
