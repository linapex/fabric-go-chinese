
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455963591577600>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package testdata

import (
	goo "github.com/hyperledger/fabric/common/metrics"
)

//这些变量应该被发现为有效的度量选项。
//即使使用了命名导入。

var (
	NamedCounter = goo.CounterOpts{
		Namespace:    "namespace",
		Subsystem:    "counter",
		Name:         "name",
		Help:         "This is some help text",
		LabelNames:   []string{"label_one", "label_two"},
		StatsdFormat: "%{#fqname}.%{label_one}.%{label_two}",
	}

	NamedGauge = goo.GaugeOpts{
		Namespace:    "namespace",
		Subsystem:    "gauge",
		Name:         "name",
		Help:         "This is some help text",
		LabelNames:   []string{"label_one", "label_two"},
		StatsdFormat: "%{#fqname}.%{label_one}.%{label_two}",
	}

	NamedHistogram = goo.HistogramOpts{
		Namespace:    "namespace",
		Subsystem:    "histogram",
		Name:         "name",
		Help:         "This is some help text",
		LabelNames:   []string{"label_one", "label_two"},
		StatsdFormat: "%{#fqname}.%{label_one}.%{label_two}",
	}
)

