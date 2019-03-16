
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:28</date>
//</624456088762191872>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package broadcast_test

import (
	"testing"

	"github.com/hyperledger/fabric/common/metrics"
	ab "github.com/hyperledger/fabric/protos/orderer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go：生成伪造者-o mock/ab_server.go--forke name ab server。弃权者
type abServer interface {
	ab.AtomicBroadcast_BroadcastServer
}

//go：生成伪造者-o mock/metrics_histogram.go——伪造名称metrics histogram。测量记录
type metricsHistogram interface {
	metrics.Histogram
}

//go：生成伪造者-o mock/metrics_counter.go——伪造名称metrics counter。计量计数器
type metricsCounter interface {
	metrics.Counter
}

//go：生成伪造者-o mock/metrics_provider.go-伪造名称metricsProvider。度量提供者
type metricsProvider interface {
	metrics.Provider
}

func TestBroadcast(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Broadcast Suite")
}

