
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456037683957760>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package operations

import (
	"sync"

	"github.com/hyperledger/fabric/common/metrics"
	"github.com/hyperledger/fabric/common/metrics/prometheus"
)

var (
	fabricVersion = metrics.GaugeOpts{
		Name:         "fabric_version",
		Help:         "The active version of Fabric.",
		LabelNames:   []string{"version"},
		StatsdFormat: "%{#fqname}.%{version}",
	}

	gaugeLock        sync.Mutex
	promVersionGauge metrics.Gauge
)

func versionGauge(provider metrics.Provider) metrics.Gauge {
	switch provider.(type) {
	case *prometheus.Provider:
		gaugeLock.Lock()
		defer gaugeLock.Unlock()
		if promVersionGauge == nil {
			promVersionGauge = provider.NewGauge(fabricVersion)
		}
		return promVersionGauge

	default:
		return provider.NewGauge(fabricVersion)
	}
}

