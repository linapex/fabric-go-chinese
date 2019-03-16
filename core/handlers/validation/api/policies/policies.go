
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:10</date>
//</624456012929175552>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package validation

import (
	"github.com/hyperledger/fabric/core/handlers/validation/api"
	"github.com/hyperledger/fabric/protos/common"
)

//政策评估者评估政策
type PolicyEvaluator interface {
	validation.Dependency

//Evaluate获取一组SignedData，并评估该组签名是否满足
//具有给定字节的策略
	Evaluate(policyBytes []byte, signatureSet []*common.SignedData) error
}

//序列化策略定义序列化策略
type SerializedPolicy interface {
	validation.ContextDatum

//bytes返回序列化策略的字节数
	Bytes() []byte
}

