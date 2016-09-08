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
	"github.com/jimmidyson/minishift/pkg/minikube/cluster"
	"github.com/jimmidyson/minishift/pkg/minikube/constants"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes a local OpenShift cluster.",
	Long: `Deletes a local OpenShift cluster. This command deletes the VM, and removes all
associated files.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting local OpenShift cluster...")
		api := libmachine.NewClient(constants.Minipath, constants.MakeMiniPath("certs"))
		defer api.Close()

		if err := cluster.DeleteHost(api); err != nil {
			fmt.Println("Errors occurred deleting machine: ", err)
			os.Exit(1)
		}
		fmt.Println("Machine deleted.")
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
