
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456035813298176>

/*
版权所有IBM公司。保留所有权利。
SPDX许可证标识符：Apache-2.0
**/


package couchdb

import (
	"time"

	"github.com/hyperledger/fabric/common/metrics"
)

var (
	apiProcessingTimeOpts = metrics.HistogramOpts{
		Namespace:    "couchdb",
		Subsystem:    "",
		Name:         "processing_time",
		Help:         "Time taken in seconds for the function to complete request to CouchDB",
		LabelNames:   []string{"database", "function_name", "result"},
		StatsdFormat: "%{#fqname}.%{database}.%{function_name}.%{result}",
	}
)

type stats struct {
	apiProcessingTime metrics.Histogram
}

func newStats(metricsProvider metrics.Provider) *stats {
	return &stats{
		apiProcessingTime: metricsProvider.NewHistogram(apiProcessingTimeOpts),
	}
}

func (s *stats) observeProcessingTime(startTime time.Time, dbName, functionName, result string) {
	s.apiProcessingTime.With(
		"database", dbName,
		"function_name", functionName,
		"result", result,
	).Observe(time.Since(startTime).Seconds())
}

