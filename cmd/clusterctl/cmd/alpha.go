/*
Copyright 2020 The Kubernetes Authors.

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
	"github.com/spf13/cobra"
)

var alphaCmd = &cobra.Command{
	Use:   "alpha",
	Short: "Commands for features in alpha",
	Long:  `These commands correspond to alpha features in clusterctl.`,
}

func init() {
	// Alpha commands should be added here.
	alphaCmd.AddCommand(rolloutCmd)
	alphaCmd.AddCommand(topologyCmd)

	RootCmd.AddCommand(alphaCmd)
}
