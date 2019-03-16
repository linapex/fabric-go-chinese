
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:27</date>
//</624456087923331072>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package blockcutter

import "github.com/hyperledger/fabric/common/metrics"

var (
	blockFillDuration = metrics.HistogramOpts{
		Namespace:    "blockcutter",
		Name:         "block_fill_duration",
		Help:         "The time from first transaction enqueing to the block being cut in seconds.",
		LabelNames:   []string{"channel"},
		StatsdFormat: "%{#fqname}.%{channel}",
	}
)

type Metrics struct {
	BlockFillDuration metrics.Histogram
}

func NewMetrics(p metrics.Provider) *Metrics {
	return &Metrics{
		BlockFillDuration: p.NewHistogram(blockFillDuration),
	}
}

