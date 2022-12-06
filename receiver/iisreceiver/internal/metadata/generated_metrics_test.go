// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
)

func TestDefaultMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	mb := NewMetricsBuilder(DefaultMetricsSettings(), component.BuildInfo{}, WithStartTime(start))
	enabledMetrics := make(map[string]bool)

	enabledMetrics["iis.connection.active"] = true
	mb.RecordIisConnectionActiveDataPoint(ts, 1)

	enabledMetrics["iis.connection.anonymous"] = true
	mb.RecordIisConnectionAnonymousDataPoint(ts, 1)

	enabledMetrics["iis.connection.attempt.count"] = true
	mb.RecordIisConnectionAttemptCountDataPoint(ts, 1)

	enabledMetrics["iis.network.blocked"] = true
	mb.RecordIisNetworkBlockedDataPoint(ts, 1)

	enabledMetrics["iis.network.file.count"] = true
	mb.RecordIisNetworkFileCountDataPoint(ts, 1, AttributeDirection(1))

	enabledMetrics["iis.network.io"] = true
	mb.RecordIisNetworkIoDataPoint(ts, 1, AttributeDirection(1))

	enabledMetrics["iis.request.count"] = true
	mb.RecordIisRequestCountDataPoint(ts, 1, AttributeRequest(1))

	enabledMetrics["iis.request.queue.age.max"] = true
	mb.RecordIisRequestQueueAgeMaxDataPoint(ts, 1)

	enabledMetrics["iis.request.queue.count"] = true
	mb.RecordIisRequestQueueCountDataPoint(ts, 1)

	enabledMetrics["iis.request.rejected"] = true
	mb.RecordIisRequestRejectedDataPoint(ts, 1)

	enabledMetrics["iis.thread.active"] = true
	mb.RecordIisThreadActiveDataPoint(ts, 1)

	enabledMetrics["iis.uptime"] = true
	mb.RecordIisUptimeDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	sm := metrics.ResourceMetrics().At(0).ScopeMetrics()
	assert.Equal(t, 1, sm.Len())
	ms := sm.At(0).Metrics()
	assert.Equal(t, len(enabledMetrics), ms.Len())
	seenMetrics := make(map[string]bool)
	for i := 0; i < ms.Len(); i++ {
		assert.True(t, enabledMetrics[ms.At(i).Name()])
		seenMetrics[ms.At(i).Name()] = true
	}
	assert.Equal(t, len(enabledMetrics), len(seenMetrics))
}

func TestAllMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	settings := MetricsSettings{
		IisConnectionActive:       MetricSettings{Enabled: true},
		IisConnectionAnonymous:    MetricSettings{Enabled: true},
		IisConnectionAttemptCount: MetricSettings{Enabled: true},
		IisNetworkBlocked:         MetricSettings{Enabled: true},
		IisNetworkFileCount:       MetricSettings{Enabled: true},
		IisNetworkIo:              MetricSettings{Enabled: true},
		IisRequestCount:           MetricSettings{Enabled: true},
		IisRequestQueueAgeMax:     MetricSettings{Enabled: true},
		IisRequestQueueCount:      MetricSettings{Enabled: true},
		IisRequestRejected:        MetricSettings{Enabled: true},
		IisThreadActive:           MetricSettings{Enabled: true},
		IisUptime:                 MetricSettings{Enabled: true},
	}
	mb := NewMetricsBuilder(settings, component.BuildInfo{}, WithStartTime(start))

	mb.RecordIisConnectionActiveDataPoint(ts, 1)
	mb.RecordIisConnectionAnonymousDataPoint(ts, 1)
	mb.RecordIisConnectionAttemptCountDataPoint(ts, 1)
	mb.RecordIisNetworkBlockedDataPoint(ts, 1)
	mb.RecordIisNetworkFileCountDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordIisNetworkIoDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordIisRequestCountDataPoint(ts, 1, AttributeRequest(1))
	mb.RecordIisRequestQueueAgeMaxDataPoint(ts, 1)
	mb.RecordIisRequestQueueCountDataPoint(ts, 1)
	mb.RecordIisRequestRejectedDataPoint(ts, 1)
	mb.RecordIisThreadActiveDataPoint(ts, 1)
	mb.RecordIisUptimeDataPoint(ts, 1)

	metrics := mb.Emit(WithIisApplicationPool("attr-val"), WithIisSite("attr-val"))

	assert.Equal(t, 1, metrics.ResourceMetrics().Len())
	rm := metrics.ResourceMetrics().At(0)
	attrCount := 0
	attrCount++
	attrVal, ok := rm.Resource().Attributes().Get("iis.application_pool")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	attrCount++
	attrVal, ok = rm.Resource().Attributes().Get("iis.site")
	assert.True(t, ok)
	assert.EqualValues(t, "attr-val", attrVal.Str())
	assert.Equal(t, attrCount, rm.Resource().Attributes().Len())

	assert.Equal(t, 1, rm.ScopeMetrics().Len())
	ms := rm.ScopeMetrics().At(0).Metrics()
	allMetricsCount := reflect.TypeOf(MetricsSettings{}).NumField()
	assert.Equal(t, allMetricsCount, ms.Len())
	validatedMetrics := make(map[string]struct{})
	for i := 0; i < ms.Len(); i++ {
		switch ms.At(i).Name() {
		case "iis.connection.active":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of active connections.", ms.At(i).Description())
			assert.Equal(t, "{connections}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["iis.connection.active"] = struct{}{}
		case "iis.connection.anonymous":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of connections established anonymously.", ms.At(i).Description())
			assert.Equal(t, "{connections}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["iis.connection.anonymous"] = struct{}{}
		case "iis.connection.attempt.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of attempts to connect to the server.", ms.At(i).Description())
			assert.Equal(t, "{attempts}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["iis.connection.attempt.count"] = struct{}{}
		case "iis.network.blocked":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of bytes blocked due to bandwidth throttling.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["iis.network.blocked"] = struct{}{}
		case "iis.network.file.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Number of transmitted files.", ms.At(i).Description())
			assert.Equal(t, "{files}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "sent", attrVal.Str())
			validatedMetrics["iis.network.file.count"] = struct{}{}
		case "iis.network.io":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total amount of bytes sent and received.", ms.At(i).Description())
			assert.Equal(t, "By", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("direction")
			assert.True(t, ok)
			assert.Equal(t, "sent", attrVal.Str())
			validatedMetrics["iis.network.io"] = struct{}{}
		case "iis.request.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of requests of a given type.", ms.At(i).Description())
			assert.Equal(t, "{requests}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			attrVal, ok := dp.Attributes().Get("request")
			assert.True(t, ok)
			assert.Equal(t, "delete", attrVal.Str())
			validatedMetrics["iis.request.count"] = struct{}{}
		case "iis.request.queue.age.max":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "Age of oldest request in the queue.", ms.At(i).Description())
			assert.Equal(t, "ms", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["iis.request.queue.age.max"] = struct{}{}
		case "iis.request.queue.count":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Current number of requests in the queue.", ms.At(i).Description())
			assert.Equal(t, "{requests}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["iis.request.queue.count"] = struct{}{}
		case "iis.request.rejected":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Total number of requests rejected.", ms.At(i).Description())
			assert.Equal(t, "{requests}", ms.At(i).Unit())
			assert.Equal(t, true, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["iis.request.rejected"] = struct{}{}
		case "iis.thread.active":
			assert.Equal(t, pmetric.MetricTypeSum, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Sum().DataPoints().Len())
			assert.Equal(t, "Current number of active threads.", ms.At(i).Description())
			assert.Equal(t, "{threads}", ms.At(i).Unit())
			assert.Equal(t, false, ms.At(i).Sum().IsMonotonic())
			assert.Equal(t, pmetric.AggregationTemporalityCumulative, ms.At(i).Sum().AggregationTemporality())
			dp := ms.At(i).Sum().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["iis.thread.active"] = struct{}{}
		case "iis.uptime":
			assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
			assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
			assert.Equal(t, "The amount of time the server has been up.", ms.At(i).Description())
			assert.Equal(t, "s", ms.At(i).Unit())
			dp := ms.At(i).Gauge().DataPoints().At(0)
			assert.Equal(t, start, dp.StartTimestamp())
			assert.Equal(t, ts, dp.Timestamp())
			assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
			assert.Equal(t, int64(1), dp.IntValue())
			validatedMetrics["iis.uptime"] = struct{}{}
		}
	}
	assert.Equal(t, allMetricsCount, len(validatedMetrics))
}

func TestNoMetrics(t *testing.T) {
	start := pcommon.Timestamp(1_000_000_000)
	ts := pcommon.Timestamp(1_000_001_000)
	settings := MetricsSettings{
		IisConnectionActive:       MetricSettings{Enabled: false},
		IisConnectionAnonymous:    MetricSettings{Enabled: false},
		IisConnectionAttemptCount: MetricSettings{Enabled: false},
		IisNetworkBlocked:         MetricSettings{Enabled: false},
		IisNetworkFileCount:       MetricSettings{Enabled: false},
		IisNetworkIo:              MetricSettings{Enabled: false},
		IisRequestCount:           MetricSettings{Enabled: false},
		IisRequestQueueAgeMax:     MetricSettings{Enabled: false},
		IisRequestQueueCount:      MetricSettings{Enabled: false},
		IisRequestRejected:        MetricSettings{Enabled: false},
		IisThreadActive:           MetricSettings{Enabled: false},
		IisUptime:                 MetricSettings{Enabled: false},
	}
	mb := NewMetricsBuilder(settings, component.BuildInfo{}, WithStartTime(start))
	mb.RecordIisConnectionActiveDataPoint(ts, 1)
	mb.RecordIisConnectionAnonymousDataPoint(ts, 1)
	mb.RecordIisConnectionAttemptCountDataPoint(ts, 1)
	mb.RecordIisNetworkBlockedDataPoint(ts, 1)
	mb.RecordIisNetworkFileCountDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordIisNetworkIoDataPoint(ts, 1, AttributeDirection(1))
	mb.RecordIisRequestCountDataPoint(ts, 1, AttributeRequest(1))
	mb.RecordIisRequestQueueAgeMaxDataPoint(ts, 1)
	mb.RecordIisRequestQueueCountDataPoint(ts, 1)
	mb.RecordIisRequestRejectedDataPoint(ts, 1)
	mb.RecordIisThreadActiveDataPoint(ts, 1)
	mb.RecordIisUptimeDataPoint(ts, 1)

	metrics := mb.Emit()

	assert.Equal(t, 0, metrics.ResourceMetrics().Len())
}
