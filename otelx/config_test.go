// Copyright © 2023 Ory Corp
// SPDX-License-Identifier: Apache-2.0

package otelx

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/sjson"

	"github.com/ory/jsonschema/v3"
)

const rootSchema = `{
  "properties": {
    "tracing": {
      "$ref": "%s"
    }
  }
}
`

func TestConfigSchema(t *testing.T) {
	t.Run("func=AddConfigSchema", func(t *testing.T) {
		c := jsonschema.NewCompiler()
		require.NoError(t, AddConfigSchema(c))

		testCases := map[string]Config{
			"jaeger with local agent": {
				ServiceName: "Ory X",
				Provider:    "jaeger",
				Providers: ProvidersConfig{
					Jaeger: JaegerConfig{
						LocalAgentAddress: "localhost:6831",
						Sampling: JaegerSampling{
							ServerURL:    "http://localhost:5778/sampling",
							TraceIdRatio: 1,
						},
					},
				},
			},
			"jaeger with collector": {
				ServiceName: "Ory X",
				Provider:    "jaeger",
				Providers: ProvidersConfig{
					Jaeger: JaegerConfig{
						CollectorScheme:  "http",
						CollectorAddress: "localhost:14268",
						Sampling: JaegerSampling{
							ServerURL:    "http://localhost:5778/sampling",
							TraceIdRatio: 1,
						},
					},
				},
			},
			"no providers": {
				ServiceName: "Ory X",
				Provider:    "noop",
			},
		}

		for name, conf := range testCases {
			t.Run(name, func(t *testing.T) {
				rawConfig, err := sjson.Set("{}", "otelx", &conf)
				require.NoError(t, err)

				require.NoError(t, c.AddResource("config", bytes.NewBufferString(fmt.Sprintf(rootSchema, ConfigSchemaID))))
				schema, err := c.Compile(context.Background(), "config")
				require.NoError(t, err)

				assert.NoError(t, schema.Validate(bytes.NewBufferString(rawConfig)))
			})
		}
	})
}
