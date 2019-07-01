/**
 * Copyright 2017 ~ 2025 the original author or authors[983708408@qq.com].
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package indicators

import (
	"github.com/gogo/protobuf/proto"
	"regexp"
	"strings"
	"time"
	"umc-agent/pkg/common"
	"umc-agent/pkg/config"
)

// Metric aggregate wrapper.
type MetricAggregator struct {
	metricType string // Temporary
	MetricAggregate
}

// Metric wrapper.
type MetricWrapper struct {
	Metric
}

// Internal create metric.
func internalNewMetric(metricName string, value float64) *MetricWrapper {
	var _metric = new(MetricWrapper)
	_metric.Tags = make(map[string]string)
	_metric.Metric.Metric = metricName
	_metric.Value = value
	return _metric
}

// Metric tags appender.
func (self *MetricWrapper) ATag(key string, value string) *MetricWrapper {
	self.Tags[key] = value
	return self
}

// New metric aggregator. Indicator names must be consistent with the global configuration,
// pay attention to case-sensitive (e.g., optional values: Redis, Kafka, Emq...),
// See: `./pkg/config/indicator_config.go#type[IndicatorProperties]` for names of members
func NewMetricAggregator(metricType string) *MetricAggregator {
	aggregator := new(MetricAggregator)
	aggregator.metricType = metricType
	aggregator.Instance = config.LocalHardwareAddrId
	aggregator.Namespace = config.GlobalConfig.Indicator.Namespace
	aggregator.Timestamp = time.Now().Unix()
	return aggregator
}

// Create metrics with the creator.
func (self *MetricAggregator) NewMetric(metricName string, value float64) *MetricWrapper {
	var (
		// Check necessary.(Note: that the project configuration
		// structure must correspond to this.)
		filters = config.GetConfig(config.IndicatorsFiledName, self.metricType, config.MetricExcludeRegexFieldName)

		// Create metricName.
		_metric = internalNewMetric(metricName, value)
	)

	// Enabled only if the configuration metricName name is included.
	reg := strings.Split(filters.ToString(), ",")
	for _, r := range reg {
		if r == "" {
			break
		}
		b, _ := (regexp.Match(r, []byte(metricName)))
		if b {
			return _metric
		}
	}

	self.Metrics = append(self.Metrics, &_metric.Metric)
	/*if common.StringsContains(reg, metricName) {
		self.Metrics = append(self.Metrics, &_metric.Metric)
	}*/
	return _metric
}

// To metric aggregate json string.
func (self *MetricAggregator) ToJSONString() string {
	return common.ToJSONString(self)
}

// To metric aggregate proto buffer.
func (self *MetricAggregator) ToProtoBuf() ([]byte, error) {
	data, err := proto.Marshal(self)
	return data, err
}

// To metric aggregate proto buffer array.
func (self *MetricAggregator) ToProtoBufArray() []byte {
	data, _ := self.ToProtoBuf()
	return data
}
