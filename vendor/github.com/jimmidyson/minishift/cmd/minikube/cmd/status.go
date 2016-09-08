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
	"fmt"
	"os"

	"github.com/docker/machine/libmachine"
	"github.com/golang/glog"
	"github.com/jimmidyson/minishift/pkg/minikube/cluster"
	"github.com/jimmidyson/minishift/pkg/minikube/constants"
	"github.com/spf13/cobra"
)

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gets the status of a local OpenShift cluster.",
	Long:  `Gets the status of a local OpenShift cluster.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := libmachine.NewClient(constants.Minipath, constants.MakeMiniPath("certs"))
		defer api.Close()
		s, err := cluster.GetHostStatus(api)
		if err != nil {
			glog.Errorln("Error getting machine status:", err)
			os.Exit(1)
		}
		fmt.Fprintln(os.Stdout, s)
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
