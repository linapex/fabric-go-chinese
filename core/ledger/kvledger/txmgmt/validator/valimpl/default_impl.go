
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456028997554176>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package valimpl

import (
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/privacyenabledstate"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/txmgr"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/validator"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/validator/internal"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/validator/statebasedval"
	"github.com/hyperledger/fabric/core/ledger/util"
	"github.com/hyperledger/fabric/protos/common"
)

var logger = flogging.MustGetLogger("valimpl")

//DefaultImpl实现接口验证程序。验证程序
//这将执行独立于特定验证方案的常见任务
//对于公共rwset的实际验证，它包含一个内部验证程序（实现接口
//internal.internal validator）例如基于状态的验证程序
type DefaultImpl struct {
	txmgr             txmgr.TxMgr
	db                privacyenabledstate.DB
	internalValidator internal.Validator
}

//newstatebasedvalidator构造一个内部管理基于状态的验证器的验证器，此外
//处理与特定验证方案无关的任务，例如分析块和处理pvt数据
func NewStatebasedValidator(txmgr txmgr.TxMgr, db privacyenabledstate.DB) validator.Validator {
	return &DefaultImpl{txmgr, db, statebasedval.NewValidator(db)}
}

//验证和批处理实现接口验证程序中的函数。
func (impl *DefaultImpl) ValidateAndPrepareBatch(blockAndPvtdata *ledger.BlockAndPvtData,
	doMVCCValidation bool) (*privacyenabledstate.UpdateBatch, []*txmgr.TxStatInfo, error) {
	block := blockAndPvtdata.Block
	logger.Debugf("ValidateAndPrepareBatch() for block number = [%d]", block.Header.Number)
	var internalBlock *internal.Block
	var txsStatInfo []*txmgr.TxStatInfo
	var pubAndHashUpdates *internal.PubAndHashUpdates
	var pvtUpdates *privacyenabledstate.PvtUpdateBatch
	var err error

	logger.Debug("preprocessing ProtoBlock...")
	if internalBlock, txsStatInfo, err = preprocessProtoBlock(impl.txmgr, impl.db.ValidateKeyValue, block, doMVCCValidation); err != nil {
		return nil, nil, err
	}

	if pubAndHashUpdates, err = impl.internalValidator.ValidateAndPrepareBatch(internalBlock, doMVCCValidation); err != nil {
		return nil, nil, err
	}
	logger.Debug("validating rwset...")
	if pvtUpdates, err = validateAndPreparePvtBatch(internalBlock, impl.db, pubAndHashUpdates, blockAndPvtdata.PvtData); err != nil {
		return nil, nil, err
	}
	logger.Debug("postprocessing ProtoBlock...")
	postprocessProtoBlock(block, internalBlock)
	logger.Debug("ValidateAndPrepareBatch() complete")

	txsFilter := util.TxValidationFlags(block.Metadata.Metadata[common.BlockMetadataIndex_TRANSACTIONS_FILTER])
	for i := range txsFilter {
		txsStatInfo[i].ValidationCode = txsFilter.Flag(i)
	}
	return &privacyenabledstate.UpdateBatch{
		PubUpdates:  pubAndHashUpdates.PubUpdates,
		HashUpdates: pubAndHashUpdates.HashUpdates,
		PvtUpdates:  pvtUpdates,
	}, txsStatInfo, nil
}

