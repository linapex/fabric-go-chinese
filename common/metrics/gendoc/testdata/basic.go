
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455963457359872>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package testdata

import (
	"time"

	"github.com/hyperledger/fabric/common/metrics"
)

//

var (
	Counter = metrics.CounterOpts{
		Namespace:    "fixtures",
		Name:         "counter",
		Help:         "This is some help text that is more than a few words long. It really can be quite long. Really long.",
		LabelNames:   []string{"label_one", "label_two"},
		StatsdFormat: "%{#fqname}.%{label_one}.%{label_two}",
	}

	Gauge = metrics.GaugeOpts{
		Namespace:    "fixtures",
		Name:         "gauge",
		Help:         "This is some help text",
		LabelNames:   []string{"label_one", "label_two"},
		StatsdFormat: "%{#fqname}.%{label_one}.%{label_two}",
	}

	Histogram = metrics.HistogramOpts{
		Namespace:    "fixtures",
		Name:         "histogram",
		Help:         "This is some help text",
		LabelNames:   []string{"label_one", "label_two"},
		StatsdFormat: "%{#fqname}.%{label_one}.%{label_two}",
	}

	ignoredStruct = struct{}{}

	ignoredInt = 0

	ignoredTime = time.Now()
)

