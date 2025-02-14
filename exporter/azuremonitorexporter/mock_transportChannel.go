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

// Code generated by mockery v1.0.0. DO NOT EDIT.

package azuremonitorexporter

import (
	contracts "github.com/microsoft/ApplicationInsights-Go/appinsights/contracts"
	mock "github.com/stretchr/testify/mock"
)

// mockTransportChannel is an autogenerated mock type for the transportChannel type
type mockTransportChannel struct {
	mock.Mock
}

// Send provides a mock function with given fields: _a0
func (_m *mockTransportChannel) Send(_a0 *contracts.Envelope) {
	_m.Called(_a0)
}
