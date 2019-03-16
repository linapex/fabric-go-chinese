
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:12</date>
//</624456021275840512>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package privacyenabledstate_test

import (
	"testing"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb/statecouchdb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb/stateleveldb"
	"github.com/hyperledger/fabric/core/ledger/mock"
	. "github.com/onsi/gomega"
)

func TestHealthCheckRegister(t *testing.T) {
	gt := NewGomegaWithT(t)
	fakeHealthCheckRegistry := &mock.HealthCheckRegistry{}

	dbProvider := &privacyenabledstate.CommonStorageDBProvider{
		VersionedDBProvider: &stateleveldb.VersionedDBProvider{},
		HealthCheckRegistry: fakeHealthCheckRegistry,
	}

	err := dbProvider.RegisterHealthChecker()
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Expect(fakeHealthCheckRegistry.RegisterCheckerCallCount()).To(Equal(0))

	dbProvider.VersionedDBProvider = &statecouchdb.VersionedDBProvider{}
	err = dbProvider.RegisterHealthChecker()
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Expect(fakeHealthCheckRegistry.RegisterCheckerCallCount()).To(Equal(1))

	arg1, arg2 := fakeHealthCheckRegistry.RegisterCheckerArgsForCall(0)
	gt.Expect(arg1).To(Equal("couchdb"))
	gt.Expect(arg2).NotTo(Equal(nil))
}

