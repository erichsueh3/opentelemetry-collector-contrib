// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for iisreceiver metrics.
type MetricsConfig struct {
	IisConnectionActive       MetricConfig `mapstructure:"iis.connection.active"`
	IisConnectionAnonymous    MetricConfig `mapstructure:"iis.connection.anonymous"`
	IisConnectionAttemptCount MetricConfig `mapstructure:"iis.connection.attempt.count"`
	IisNetworkBlocked         MetricConfig `mapstructure:"iis.network.blocked"`
	IisNetworkFileCount       MetricConfig `mapstructure:"iis.network.file.count"`
	IisNetworkIo              MetricConfig `mapstructure:"iis.network.io"`
	IisRequestCount           MetricConfig `mapstructure:"iis.request.count"`
	IisRequestQueueAgeMax     MetricConfig `mapstructure:"iis.request.queue.age.max"`
	IisRequestQueueCount      MetricConfig `mapstructure:"iis.request.queue.count"`
	IisRequestRejected        MetricConfig `mapstructure:"iis.request.rejected"`
	IisThreadActive           MetricConfig `mapstructure:"iis.thread.active"`
	IisUptime                 MetricConfig `mapstructure:"iis.uptime"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		IisConnectionActive: MetricConfig{
			Enabled: true,
		},
		IisConnectionAnonymous: MetricConfig{
			Enabled: true,
		},
		IisConnectionAttemptCount: MetricConfig{
			Enabled: true,
		},
		IisNetworkBlocked: MetricConfig{
			Enabled: true,
		},
		IisNetworkFileCount: MetricConfig{
			Enabled: true,
		},
		IisNetworkIo: MetricConfig{
			Enabled: true,
		},
		IisRequestCount: MetricConfig{
			Enabled: true,
		},
		IisRequestQueueAgeMax: MetricConfig{
			Enabled: true,
		},
		IisRequestQueueCount: MetricConfig{
			Enabled: true,
		},
		IisRequestRejected: MetricConfig{
			Enabled: true,
		},
		IisThreadActive: MetricConfig{
			Enabled: true,
		},
		IisUptime: MetricConfig{
			Enabled: true,
		},
	}
}

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`
}

// ResourceAttributesConfig provides config for iisreceiver resource attributes.
type ResourceAttributesConfig struct {
	IisApplicationPool ResourceAttributeConfig `mapstructure:"iis.application_pool"`
	IisSite            ResourceAttributeConfig `mapstructure:"iis.site"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		IisApplicationPool: ResourceAttributeConfig{
			Enabled: true,
		},
		IisSite: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}

// MetricsBuilderConfig is a configuration for iisreceiver metrics builder.
type MetricsBuilderConfig struct {
	Metrics            MetricsConfig            `mapstructure:"metrics"`
	ResourceAttributes ResourceAttributesConfig `mapstructure:"resource_attributes"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics:            DefaultMetricsConfig(),
		ResourceAttributes: DefaultResourceAttributesConfig(),
	}
}
