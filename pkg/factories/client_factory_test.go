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

package factories

import (
	"testing"

	"gotest.tools/assert"

	"github.com/maximilien/kn-source-pkg/pkg/types"
)

func TestNewDefaultClientFactory(t *testing.T) {
	clientFactory := NewDefaultClientFactory(&types.KnSourceParams{})

	assert.Assert(t, clientFactory != nil)
}

func TestClientFactory_KnSourceParams(t *testing.T) {
	knSourceParams := &types.KnSourceParams{}
	clientFactory := NewDefaultClientFactory(knSourceParams)

	assert.Equal(t, clientFactory.KnSourceParams(), knSourceParams)
}

func TestCreateKnSourceClient(t *testing.T) {
	clientFactory := NewDefaultClientFactory(&types.KnSourceParams{})
	client := clientFactory.CreateKnSourceClient("fake-namespace")

	assert.Assert(t, client != nil)
	assert.Equal(t, client.Namespace(), "fake-namespace")
}