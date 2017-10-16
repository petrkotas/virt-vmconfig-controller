// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// RootCmd call controller without any params, it will start with defaults
var RootCmd = &cobra.Command{
	Use:   "virt-vmconfig-controller",
	Short: "Starts the controller",
	Long: `Starts the virt-VM-config controller.
	It watches for the change from state: stopped to
	state: started.`,

	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Called the virt-controller")
	},
}

// Run starts the whole app
func Run() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal("Somethig very vad happened")
	}
}
