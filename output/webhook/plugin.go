// Copyright 2023 Blink Labs, LLC.
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

package webhook

import (
	"github.com/blinklabs-io/snek/plugin"
)

var cmdlineOptions struct {
	url string
}

func init() {
	plugin.Register(
		plugin.PluginEntry{
			Type:               plugin.PluginTypeOutput,
			Name:               "webhook",
			Description:        "send events via HTTP POST to a webhook server",
			NewFromOptionsFunc: NewFromCmdlineOptions,
			Options: []plugin.PluginOption{
				{
					Name:         "url",
					Type:         plugin.PluginOptionTypeString,
					Description:  "specifies the url to use",
					DefaultValue: "http://localhost:3000",
					Dest:         &(cmdlineOptions.url),
				},
			},
		},
	)
}

func NewFromCmdlineOptions() plugin.Plugin {
	p := New(
		WithUrl(cmdlineOptions.url),
	)
	return p
}