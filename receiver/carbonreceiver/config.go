// Copyright 2019, OpenTelemetry Authors
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

package carbonreceiver

import (
	"time"

	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/config/confignet"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/carbonreceiver/protocol"
)

const (
	// parserConfigSection is the name that must be used for the parser settings
	// in the configuration struct. The metadata mapstructure for the parser
	// should use the same string.
	parserConfigSection = "parser"
)

// Config defines configuration for the Carbon receiver.
type Config struct {
	config.ReceiverSettings `mapstructure:",squash"`

	confignet.NetAddr `mapstructure:",squash"`

	// TCPIdleTimeout is the timout for idle TCP connections, it is ignored
	// if transport being used is UDP.
	TCPIdleTimeout time.Duration `mapstructure:"tcp_idle_timeout"`

	// Parser specifies a parser and the respective configuration to be used
	// by the receiver.
	Parser *protocol.Config `mapstructure:"parser"`
}
