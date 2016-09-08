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
	goflag "flag"
	"io/ioutil"
	"os"

	"github.com/docker/machine/libmachine/log"
	"github.com/golang/glog"
	"github.com/jimmidyson/minishift/pkg/minikube/config"
	"github.com/jimmidyson/minishift/pkg/minikube/constants"
	"github.com/jimmidyson/minishift/pkg/minikube/update"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var dirs = [...]string{
	constants.Minipath,
	constants.MakeMiniPath("certs"),
	constants.MakeMiniPath("machines"),
	constants.MakeMiniPath("cache"),
	constants.MakeMiniPath("cache", "iso"),
	constants.MakeMiniPath("config")}

const (
	showLibmachineLogs = "show-libmachine-logs"
)

var (
	enableUpdateNotification = true
)

var viperWhiteList = []string{
	"v",
	"alsologtostderr",
	"log_dir",
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "minishift",
	Short: "Minishift is a tool for managing local OpenShift clusters.",
	Long:  `Minishift is a CLI tool that provisions and manages single-node OpenShift clusters optimized for development workflows.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		for _, path := range dirs {
			if err := os.MkdirAll(path, 0777); err != nil {
				glog.Exitf("Error creating minishift directory: %s", err)
			}
		}

		shouldShowLibmachineLogs := viper.GetBool(showLibmachineLogs)

		log.SetDebug(shouldShowLibmachineLogs)
		if !shouldShowLibmachineLogs {
			log.SetOutWriter(ioutil.Discard)
			log.SetErrWriter(ioutil.Discard)
		}

		if enableUpdateNotification {
			update.MaybeUpdateFromGithub(os.Stdout)
		}
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		glog.Exitln(err)
	}
}

// Handle config values for flags used in external packages (e.g. glog)
// by setting them directly, using values from viper when not passed in as args
func setFlagsUsingViper() {
	for _, config := range viperWhiteList {
		var a = pflag.Lookup(config)
		viper.SetDefault(a.Name, a.DefValue)
		// If the flag is set, override viper value
		if a.Changed {
			viper.Set(a.Name, a.Value.String())
		}
		// Viper will give precedence first to calls to the Set command,
		// then to values from the config.yml
		a.Value.Set(viper.GetString(a.Name))
	}
}

func init() {
	RootCmd.PersistentFlags().Bool(showLibmachineLogs, false, "Whether or not to show logs from libmachine.")
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)
	viper.BindPFlags(RootCmd.PersistentFlags())
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	configPath := constants.MakeMiniPath("config")
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		glog.Warningf("Error reading config file at %s: %s", configPath, err)
	}
	setupViper()
}

func setupViper() {
	viper.SetEnvPrefix(constants.MiniShiftEnvPrefix)
	viper.AutomaticEnv()

	viper.SetDefault(config.WantUpdateNotification, true)
	viper.SetDefault(config.ReminderWaitPeriodInHours, 24)
	setFlagsUsingViper()
}
