/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package cmds

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"github.com/fabric8io/gofabric8/util"
	"github.com/jimmidyson/minishift/pkg/minikube/update"
	"github.com/spf13/cobra"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

const (
	minishiftFlag        = "minishift"
	minishiftOwner       = "jimmidyson"
	minishift            = "minishift"
	minishiftDownloadURL = "https://github.com/jimmidyson/"
	kubectl              = "kubectl"
	writeFileLocation    = "~/fabric8/bin/"
)

var (
	kubeDistroOrg   = "kubernetes"
	kubeDistroRepo  = "minikube"
	kubeDownloadURL = "https://storage.googleapis.com/"
	downloadPath    = ""
)

// NewCmdInstall installs the dependencies to run the fabric8 microservices platform
func NewCmdInstall(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Installs the dependencies to locally run the fabric8 microservices platform",
		Long:  `Installs the dependencies to locally run the fabric8 microservices platform`,

		Run: func(cmd *cobra.Command, args []string) {

			isMinishift := cmd.Flags().Lookup(minishiftFlag).Value.String() == "true"

			err := downloadKubernetes(isMinishift)
			if err != nil {
				util.Warnf("Unable to download kubernetes distro %v", err)
			}

			err = downloadClient(isMinishift)
			if err != nil {
				util.Warnf("Unable to download client %v", err)
			}

		},
	}
	cmd.PersistentFlags().Bool(minishiftFlag, false, "Install minishift rather than minikube")
	return cmd
}

func downloadKubernetes(isMinishift bool) (err error) {
	os := runtime.GOOS
	arch := runtime.GOARCH
	if isMinishift {
		kubeDistroOrg = minishiftOwner
		kubeDistroRepo = minishift
		kubeDownloadURL = minishiftDownloadURL
		downloadPath = "download/"
	}

	latestVersion, err := update.GetLatestVersionFromGitHub(kubeDistroOrg, kubeDistroRepo)
	if err != nil {
		util.Errorf("Unable to get latest version for %s/%s %v", kubeDistroOrg, kubeDistroRepo, err)
		return err
	}

	kubeURL := fmt.Sprintf(kubeDownloadURL+kubeDistroRepo+"/releases/"+downloadPath+"v%s/%s-%s-%s", latestVersion, kubeDistroRepo, os, arch)
	util.Infof("Downloading %s\n", kubeURL)

	err = downloadFile(writeFileLocation+kubeDistroRepo, kubeURL)

	return err
}

func downloadClient(isMinishift bool) (err error) {
	util.Info("Downloading client..")

	os := runtime.GOOS
	arch := runtime.GOARCH

	latestVersion, err := update.GetLatestVersionFromGitHub(kubeDistroOrg, kubeDistroRepo)
	if err != nil {
		util.Errorf("Unable to get latest version for %s/%s %v", kubeDistroOrg, kubeDistroRepo, err)
		return
	}

	if isMinishift {
		return fmt.Errorf("Openshift client download not yet supported")
	}

	clientURL := fmt.Sprintf("https://storage.googleapis.com/kubernetes-release/release/v%s/bin/%s/%s/%s", latestVersion, os, arch, kubectl)

	util.Infof("Dwonloading %s\n", clientURL)

	err = downloadFile(writeFileLocation+kubectl, clientURL)

	return err
}
func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
