
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:08</date>
//</624456005631086592>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package dockercontroller

import "github.com/hyperledger/fabric/common/metrics"

var (
	chaincodeImageBuildDuration = metrics.HistogramOpts{
		Namespace:    "dockercontroller",
		Name:         "chaincode_container_build_duration",
		Help:         "The time to build a chaincode image in seconds.",
		LabelNames:   []string{"chaincode", "success"},
		StatsdFormat: "%{#fqname}.%{chaincode}.%{success}",
	}
)

type BuildMetrics struct {
	ChaincodeImageBuildDuration metrics.Histogram
}

func NewBuildMetrics(p metrics.Provider) *BuildMetrics {
	return &BuildMetrics{
		ChaincodeImageBuildDuration: p.NewHistogram(chaincodeImageBuildDuration),
	}
}

