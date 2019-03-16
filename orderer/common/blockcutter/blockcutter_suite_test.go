
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:27</date>
//</624456087537455104>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package blockcutter_test

import (
	"testing"

	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/metrics"
	"github.com/hyperledger/fabric/orderer/common/blockcutter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go：生成伪造者-o mock/metrics_histogram.go——伪造名称metrics histogram。测量记录
type metricsHistogram interface {
	metrics.Histogram
}

//go：生成伪造者-o mock/metrics_provider.go-伪造名称metricsProvider。度量提供者
type metricsProvider interface {
	metrics.Provider
}

//go：生成仿冒者-o mock/config-fetcher.go--forke-name-orderconfigfetcher。订单控制器
type ordererConfigFetcher interface {
	blockcutter.OrdererConfigFetcher
}

//go：生成伪造者-o mock/order_config.go--forke name orderconfig。有序配置
type ordererConfig interface {
	channelconfig.Orderer
}

func TestBlockcutter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Blockcutter Suite")
}

