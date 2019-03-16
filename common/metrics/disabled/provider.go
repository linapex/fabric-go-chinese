
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455962886934528>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package disabled

import (
	"github.com/hyperledger/fabric/common/metrics"
)

type Provider struct{}

func (p *Provider) NewCounter(o metrics.CounterOpts) metrics.Counter       { return &Counter{} }
func (p *Provider) NewGauge(o metrics.GaugeOpts) metrics.Gauge             { return &Gauge{} }
func (p *Provider) NewHistogram(o metrics.HistogramOpts) metrics.Histogram { return &Histogram{} }

type Counter struct{}

func (c *Counter) Add(delta float64) {}
func (c *Counter) With(labelValues ...string) metrics.Counter {
	return c
}

type Gauge struct{}

func (g *Gauge) Add(delta float64) {}
func (g *Gauge) Set(delta float64) {}
func (g *Gauge) With(labelValues ...string) metrics.Gauge {
	return g
}

type Histogram struct{}

func (h *Histogram) Observe(value float64) {}
func (h *Histogram) With(labelValues ...string) metrics.Histogram {
	return h
}

