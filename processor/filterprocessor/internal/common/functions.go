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

package common // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/filterprocessor/internal/common"

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/internal/filter/filterottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl"
	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/ottl/contexts/ottlmetric"
)

func MetricFunctions() map[string]ottl.Factory[ottlmetric.TransformContext] {
	funcs := filterottl.StandardMetricFuncs()
	hasAttributeKeyOnDatapoint := newHasAttributeKeyOnDatapointFactory()
	funcs[hasAttributeKeyOnDatapoint.Name()] = hasAttributeKeyOnDatapoint

	hasAttributeOnDatapoint := newHasAttributeOnDatapointFactory()
	funcs[hasAttributeOnDatapoint.Name()] = hasAttributeOnDatapoint
	return funcs
}

type hasAttributeOnDatapointArguments struct {
	Key         string `ottlarg:"0"`
	ExpectedVal string `ottlarg:"1"`
}

func newHasAttributeOnDatapointFactory() ottl.Factory[ottlmetric.TransformContext] {
	return ottl.NewFactory("HasAttrOnDatapoint", &hasAttributeOnDatapointArguments{}, createHasAttributeOnDatapointFunction)
}

func createHasAttributeOnDatapointFunction(fCtx ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[ottlmetric.TransformContext], error) {
	args, ok := oArgs.(*hasAttributeOnDatapointArguments)

	if !ok {
		return nil, fmt.Errorf("hasAttributeOnDatapointFactory args must be of type *hasAttributeOnDatapointArguments")
	}

	return hasAttributeOnDatapoint(args.Key, args.ExpectedVal)
}

func hasAttributeOnDatapoint(key string, expectedVal string) (ottl.ExprFunc[ottlmetric.TransformContext], error) {
	return func(ctx context.Context, tCtx ottlmetric.TransformContext) (interface{}, error) {
		return checkDataPoints(tCtx, key, &expectedVal)
	}, nil
}

type hasAttributeKeyOnDatapointArguments struct {
	Key string `ottlarg:"0"`
}

func newHasAttributeKeyOnDatapointFactory() ottl.Factory[ottlmetric.TransformContext] {
	return ottl.NewFactory("HasAttrKeyOnDatapoint", &hasAttributeKeyOnDatapointArguments{}, createHasAttributeKeyOnDatapointFunction)
}

func createHasAttributeKeyOnDatapointFunction(fCtx ottl.FunctionContext, oArgs ottl.Arguments) (ottl.ExprFunc[ottlmetric.TransformContext], error) {
	args, ok := oArgs.(*hasAttributeOnDatapointArguments)

	if !ok {
		return nil, fmt.Errorf("hasAttributeKeyOnDatapointFactory args must be of type *hasAttributeOnDatapointArguments")
	}

	return hasAttributeKeyOnDatapoint(args.Key)
}

func hasAttributeKeyOnDatapoint(key string) (ottl.ExprFunc[ottlmetric.TransformContext], error) {
	return func(ctx context.Context, tCtx ottlmetric.TransformContext) (interface{}, error) {
		return checkDataPoints(tCtx, key, nil)
	}, nil
}

func checkDataPoints(tCtx ottlmetric.TransformContext, key string, expectedVal *string) (interface{}, error) {
	metric := tCtx.GetMetric()
	switch metric.Type() {
	case pmetric.MetricTypeSum:
		return checkNumberDataPointSlice(metric.Sum().DataPoints(), key, expectedVal), nil
	case pmetric.MetricTypeGauge:
		return checkNumberDataPointSlice(metric.Gauge().DataPoints(), key, expectedVal), nil
	case pmetric.MetricTypeHistogram:
		return checkHistogramDataPointSlice(metric.Histogram().DataPoints(), key, expectedVal), nil
	case pmetric.MetricTypeExponentialHistogram:
		return checkExponentialHistogramDataPointSlice(metric.ExponentialHistogram().DataPoints(), key, expectedVal), nil
	case pmetric.MetricTypeSummary:
		return checkSummaryDataPointSlice(metric.Summary().DataPoints(), key, expectedVal), nil
	}
	return nil, fmt.Errorf("unknown metric type")
}

func checkNumberDataPointSlice(dps pmetric.NumberDataPointSlice, key string, expectedVal *string) bool {
	for i := 0; i < dps.Len(); i++ {
		dp := dps.At(i)
		value, ok := dp.Attributes().Get(key)
		if ok {
			if expectedVal != nil {
				return value.Str() == *expectedVal
			}
			return true
		}
	}
	return false
}

func checkHistogramDataPointSlice(dps pmetric.HistogramDataPointSlice, key string, expectedVal *string) bool {
	for i := 0; i < dps.Len(); i++ {
		dp := dps.At(i)
		value, ok := dp.Attributes().Get(key)
		if ok {
			if expectedVal != nil {
				return value.Str() == *expectedVal
			}
			return true
		}
	}
	return false
}

func checkExponentialHistogramDataPointSlice(dps pmetric.ExponentialHistogramDataPointSlice, key string, expectedVal *string) bool {
	for i := 0; i < dps.Len(); i++ {
		dp := dps.At(i)
		value, ok := dp.Attributes().Get(key)
		if ok {
			if expectedVal != nil {
				return value.Str() == *expectedVal
			}
			return true
		}
	}
	return false
}

func checkSummaryDataPointSlice(dps pmetric.SummaryDataPointSlice, key string, expectedVal *string) bool {
	for i := 0; i < dps.Len(); i++ {
		dp := dps.At(i)
		value, ok := dp.Attributes().Get(key)
		if ok {
			if expectedVal != nil {
				return value.Str() == *expectedVal
			}
			return true
		}
	}
	return false
}
