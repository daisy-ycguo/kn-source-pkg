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

package main

import (
	"fmt"
	"os"

	"github.com/maximilien/kn-source-pkg/pkg/core"
	"github.com/maximilien/kn-source-pkg/pkg/factories"
)

func main() {
	knSourceParams := factories.NewDefaultParamsFactory().CreateKnSourceParams()

	clientFactory := factories.NewDefaultClientFactory(knSourceParams)
	commandFactory := factories.NewDefaultCommandFactory(knSourceParams)
	flagsFactory := factories.NewDefaultFlagsFactory(knSourceParams)
	runEFactory := factories.NewDefaultRunEFactory(knSourceParams, clientFactory)

	err := core.NewKnSourceCommand(knSourceParams, commandFactory, flagsFactory, runEFactory).Execute()
	if err != nil {
		if err.Error() != "subcommand is required" {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
