package funcs

import (
	"github.com/open-falcon/common/model"
	"strings"
)

func NewMetricValue(metric string, val interface{}, dataType string, tags ...string) *model.MetricValue {
	mv := model.MetricValue{
		Metric: metric,
		Value:  val,
		Type:   dataType,
	}

	size := len(tags)

	if size > 0 {
		mv.Tags = strings.Join(tags, ",")
	}

	return &mv
}
// 原值类型
func GaugeValue(metric string, val interface{}, tags ...string) *model.MetricValue {
	return NewMetricValue(metric, val, "GAUGE", tags...)
}

// 计数器类型
func CounterValue(metric string, val interface{}, tags ...string) *model.MetricValue {
	return NewMetricValue(metric, val, "COUNTER", tags...)
}
