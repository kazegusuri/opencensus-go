// Copyright 2018, OpenCensus Authors
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

package trace

import (
	"reflect"
	"testing"
)

func TestApplyZeroConfig(t *testing.T) {
	cfg := config.Load().(*Config)
	ApplyConfig(Config{})
	currentCfg := config.Load().(*Config)

	if got, want := reflect.ValueOf(currentCfg.DefaultSampler).Pointer(), reflect.ValueOf(cfg.DefaultSampler).Pointer(); got != want {
		t.Fatalf("config.DefaultSampler = %#v; want %#v", got, want)
	}
	if got, want := reflect.ValueOf(currentCfg.DefaultAttributes).Pointer(), reflect.ValueOf(cfg.DefaultAttributes).Pointer(); got != want {
		t.Fatalf("config.DefaultAttributes = %#v; want %#v", got, want)
	}
	if got, want := currentCfg.IDGenerator, cfg.IDGenerator; got != want {
		t.Fatalf("config.IDGenerator = %#v; want %#v", got, want)
	}
}

func TestApplyConfig(t *testing.T) {
	initCfg := config.Load().(*Config)
	defer config.Store(initCfg)

	config.Store(&Config{})
	cfg := Config{
		DefaultSampler:    ProbabilitySampler(0),
		DefaultAttributes: []Attribute{StringAttribute("foo", "bar")},
		IDGenerator:       &defaultIDGenerator{},
	}
	ApplyConfig(cfg)

	currentCfg := config.Load().(*Config)
	if got, want := reflect.ValueOf(currentCfg.DefaultSampler).Pointer(), reflect.ValueOf(cfg.DefaultSampler).Pointer(); got != want {
		t.Fatalf("config.DefaultSampler = %#v; want %#v", got, want)
	}
	if got, want := reflect.ValueOf(currentCfg.DefaultAttributes).Pointer(), reflect.ValueOf(cfg.DefaultAttributes).Pointer(); got != want {
		t.Fatalf("config.DefaultAttributes = %#v; want %#v", got, want)
	}
	if got, want := currentCfg.IDGenerator, cfg.IDGenerator; got != want {
		t.Fatalf("config.IDGenerator = %#v; want %#v", got, want)
	}
}
