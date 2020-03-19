// Copyright © 2018 The Knative Authors
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

package core

import (
	"github.com/maximilien/kn-source-pkg/pkg/commands"

	"github.com/spf13/cobra"
)

func NewKnSourceCommand(commandFactory commands.CommandFactory, runEFactory commands.RunEFactory) *cobra.Command {
	params := &commands.KnSourceParams{}

	rootCmd := &cobra.Command{
		Use:   "source",
		Short: "Knative eventing {{.Name}} source plugin",
		Long:  "Manage your Knative {{.Name}} eventing sources",

		// Disable docs header
		DisableAutoGenTag: true,

		// Affects children as well
		SilenceUsage: true,

		// Prevents Cobra from dealing with errors as we deal with them in main.go
		SilenceErrors: true,
	}

	if params.Output != nil {
		rootCmd.SetOutput(params.Output)
	}

	//TODO: add common source commands flags here

	createCmd := commandFactory.CreateCommand(params)
	createCmd.RunE = runEFactory.CreateRunE()
	rootCmd.AddCommand(createCmd)

	deleteCmd := commandFactory.DeleteCommand(params)
	deleteCmd.RunE = runEFactory.DeleteRunE()
	rootCmd.AddCommand(deleteCmd)

	updateCmd := commandFactory.UpdateCommand(params)
	updateCmd.RunE = runEFactory.UpdateRunE()
	rootCmd.AddCommand(updateCmd)

	describeCmd := commandFactory.DescribeCommand(params)
	describeCmd.RunE = runEFactory.DescribeRunE()
	rootCmd.AddCommand(describeCmd)

	// Initialize default `help` cmd early to prevent unknown command errors
	rootCmd.InitDefaultHelpCmd()

	return rootCmd
}
