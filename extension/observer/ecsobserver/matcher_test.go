// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ecsobserver

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestMatcherType(t *testing.T) {
	m := map[matcherType]string{
		matcherTypeService:        "service",
		matcherTypeTaskDefinition: "task_definition",
		matcherTypeDockerLabel:    "docker_label",
		-1:                        "unknown_matcher_type",
	}
	for k, v := range m {
		assert.Equal(t, v, k.String())
	}
}

func TestNewMatchers(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var c Config
		_, err := newMatchers(c, matcherOptions{})
		require.Error(t, err)
	})

	t.Run("invalid", func(t *testing.T) {
		c := Config{
			DockerLabels: []DockerLabelConfig{
				{
					// no port label
				},
			},
		}
		_, err := newMatchers(c, matcherOptions{})
		require.Error(t, err)
	})

	t.Run("valid", func(t *testing.T) {
		c := Config{
			DockerLabels: []DockerLabelConfig{
				{
					PortLabel: "PROM_PORT",
				},
			},
			Services: []ServiceConfig{
				{
					NamePattern: "^nginx-.*$",
				},
			},
			TaskDefinitions: []TaskDefinitionConfig{
				{
					ArnPattern: "arn:.*jmx.*",
				},
			},
		}
		m, err := newMatchers(c, matcherOptions{Logger: zap.NewExample()})
		require.NoError(t, err)
		assert.Len(t, m, 3)
	})
}

func TestMatchedContainer_MergeTargets(t *testing.T) {
	t.Run("add new targets", func(t *testing.T) {
		m := matchedContainer{
			Targets: []matchedTarget{
				{
					Port:        1234,
					MetricsPath: "/m1",
				},
				{
					Port:        1235,
					MetricsPath: "/m2",
				},
			},
		}
		newTargets := []matchedTarget{
			{
				Port:        1234,
				MetricsPath: "/not-m1", // different path
			},
			{
				Port:        1235, // different port
				MetricsPath: "/m1",
			},
		}
		m.MergeTargets(newTargets)
		assert.Len(t, m.Targets, 4)
		assert.Equal(t, m.Targets[3].MetricsPath, "/m1") // order is append
	})

	t.Run("respect existing targets", func(t *testing.T) {
		m := matchedContainer{
			Targets: []matchedTarget{
				{
					MatcherType: matcherTypeService,
					Port:        1234,
					MetricsPath: "/m1",
				},
				{
					Port:        1235,
					MetricsPath: "/m2",
				},
			},
		}
		newTargets := []matchedTarget{
			{
				MatcherType: matcherTypeDockerLabel, // different matcher
				Port:        1234,
				MetricsPath: "/m1",
			},
			{
				Port:        1235, // different port
				MetricsPath: "/m1",
			},
		}
		m.MergeTargets(newTargets)
		assert.Len(t, m.Targets, 3)
		assert.Equal(t, matcherTypeService, m.Targets[0].MatcherType)
	})
}
