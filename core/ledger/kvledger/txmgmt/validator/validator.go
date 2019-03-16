
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456028926251008>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/


package validator

import (
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/txmgr"
)

//Validator validates the transactions present in a block and returns a batch that should be used to update the state
type Validator interface {
	ValidateAndPrepareBatch(blockAndPvtdata *ledger.BlockAndPvtData, doMVCCValidation bool) (
		*privacyenabledstate.UpdateBatch, []*txmgr.TxStatInfo, error,
	)
}

//如果公共读写集中存在集合的哈希，则将引发errpvtDataHashMissMatch
//与块中提供的相应pvt数据不匹配以进行验证
type ErrPvtdataHashMissmatch struct {
	Msg string
}

func (e *ErrPvtdataHashMissmatch) Error() string {
	return e.Msg
}

