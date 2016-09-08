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

package cmd

import (
	"os"

	"github.com/docker/machine/libmachine"
	"github.com/golang/glog"
	"github.com/jimmidyson/minishift/pkg/minikube/cluster"
	"github.com/jimmidyson/minishift/pkg/minikube/constants"
	"github.com/spf13/cobra"
)

// sshCmd represents the docker-ssh command
var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Log into or run a command on a machine with SSH; similar to 'docker-machine ssh'",
	Long:  "Log into or run a command on a machine with SSH; similar to 'docker-machine ssh'",
	Run: func(cmd *cobra.Command, args []string) {
		api := libmachine.NewClient(constants.Minipath, constants.MakeMiniPath("certs"))
		defer api.Close()
		err := cluster.CreateSSHShell(api, args)
		if err != nil {
			glog.Errorln("Error attempting to ssh into machine: ", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCmd.AddCommand(sshCmd)
}
