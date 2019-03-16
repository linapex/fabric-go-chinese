
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:29</date>
//</624456093904408576>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package msgprocessor

import (
	"time"

	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/common/crypto"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/pkg/errors"
)

type resources interface {
//orderconfig返回通道的config.order
//以及医嘱者配置是否存在
	OrdererConfig() (channelconfig.Orderer, bool)
}

//NewExpirationRejectRule返回拒绝由标识签名的邮件的规则
//由于该功能处于活动状态，谁的身份已过期
func NewExpirationRejectRule(filterSupport resources) Rule {
	return &expirationRejectRule{filterSupport: filterSupport}
}

type expirationRejectRule struct {
	filterSupport resources
}

//应用检查创建信封的标识是否已过期
func (exp *expirationRejectRule) Apply(message *common.Envelope) error {
	ordererConf, ok := exp.filterSupport.OrdererConfig()
	if !ok {
		logger.Panic("Programming error: orderer config not found")
	}
	if !ordererConf.Capabilities().ExpirationCheck() {
		return nil
	}
	signedData, err := message.AsSignedData()

	if err != nil {
		return errors.Errorf("could not convert message to signedData: %s", err)
	}
	expirationTime := crypto.ExpiresAt(signedData[0].Identity)
//标识不能过期，或者标识尚未过期
	if expirationTime.IsZero() || time.Now().Before(expirationTime) {
		return nil
	}
	return errors.New("identity expired")
}

