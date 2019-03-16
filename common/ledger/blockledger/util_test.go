
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:57</date>
//</624455961200824320>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package blockledger_test

import (
	"testing"

	"github.com/hyperledger/fabric/common/deliver/mock"
	"github.com/hyperledger/fabric/common/ledger/blockledger"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/stretchr/testify/assert"
)

func TestClose(t *testing.T) {
	for _, testCase := range []struct {
		name               string
		status             common.Status
		isIteratorNil      bool
		expectedCloseCount int
	}{
		{
			name:          "nil iterator",
			isIteratorNil: true,
		},
		{
			name:               "Next() fails",
			status:             common.Status_INTERNAL_SERVER_ERROR,
			expectedCloseCount: 1,
		},
		{
			name:               "Next() succeeds",
			status:             common.Status_SUCCESS,
			expectedCloseCount: 1,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			var iterator *mock.BlockIterator
			reader := &mock.BlockReader{}
			if !testCase.isIteratorNil {
				iterator = &mock.BlockIterator{}
				iterator.NextReturns(&common.Block{}, testCase.status)
				reader.IteratorReturns(iterator, 1)
			}

			blockledger.GetBlock(reader, 1)
			if !testCase.isIteratorNil {
				assert.Equal(t, testCase.expectedCloseCount, iterator.CloseCallCount())
			}
		})
	}
}

