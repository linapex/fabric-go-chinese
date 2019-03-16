
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456035897184256>

/*
版权所有IBM公司。保留所有权利。
SPDX许可证标识符：Apache-2.0
**/


package couchdb

import (
	"context"
	"net/http"
	"net/url"
	"testing"

	"github.com/hyperledger/fabric/common/metrics/disabled"
	"github.com/hyperledger/fabric/common/metrics/metricsfakes"
	. "github.com/onsi/gomega"
)

func TestAPIProcessTimeMetric(t *testing.T) {
	gt := NewGomegaWithT(t)
	fakeHistogram := &metricsfakes.Histogram{}
	fakeHistogram.WithReturns(fakeHistogram)

//创建新的沙发实例
	couchInstance, err := CreateCouchInstance(
		couchDBDef.URL,
		couchDBDef.Username,
		couchDBDef.Password,
		0,
		couchDBDef.MaxRetriesOnStartup,
		couchDBDef.RequestTimeout,
		couchDBDef.CreateGlobalChangesDB,
		&disabled.Provider{},
	)
	gt.Expect(err).NotTo(HaveOccurred(), "Error when trying to create couch instance")

	couchInstance.stats = &stats{
		apiProcessingTime: fakeHistogram,
	}

url, err := url.Parse("http://LoaAuth: 0”
	gt.Expect(err).NotTo(HaveOccurred(), "Error when trying to parse URL")

	couchInstance.handleRequest(context.Background(), http.MethodGet, "db_name", "function_name", url, nil, "", "", 0, true, nil)
	gt.Expect(fakeHistogram.ObserveCallCount()).To(Equal(1))
	gt.Expect(fakeHistogram.ObserveArgsForCall(0)).NotTo(BeZero())
	gt.Expect(fakeHistogram.WithArgsForCall(0)).To(Equal([]string{
		"database", "db_name",
		"function_name", "function_name",
		"result", "0",
	}))
}

