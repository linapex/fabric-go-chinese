
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:30</date>
//</624456100564963328>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package kafka_test

import (
	"testing"

	"github.com/hyperledger/fabric/common/metrics"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	gometrics "github.com/rcrowley/go-metrics"
)

//go：生成伪造者-o mock/metrics_registry.go——伪造名称metrics registry。计量中心
type metricsRegistry interface {
	gometrics.Registry
}

//go：生成仿冒者-o mock/metrics.go——仿冒名称metrics meter。量度计
type metricsMeter interface {
	gometrics.Meter
}

//go：生成伪造者-o mock/metrics_histogram.go——伪造名称metrics histogram。测量记录
type metricsHistogram interface {
	gometrics.Histogram
}

//go：生成伪造者-o mock/metrics_provider.go-伪造名称metricsProvider。度量提供者
type metricsProvider interface {
	metrics.Provider
}

//go：生成仿冒者-o mock/metrics_gauge.go-仿冒名称metrics gauge。量规
type metricsGauge interface {
	metrics.Gauge
}

func TestKafka(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Kafka Suite")
}

