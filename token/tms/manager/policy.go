
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:38</date>
//</624456130562625536>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package manager

import (
	"github.com/hyperledger/fabric/token/identity"
	"github.com/pkg/errors"
)

//allissuingvalidator允许通道的所有成员颁发新令牌。
type AllIssuingValidator struct {
	Deserializer identity.Deserializer
}

//如果传递的创建者可以颁发传递类型的令牌，则validate返回no error，否则返回错误。
func (p *AllIssuingValidator) Validate(creator identity.PublicInfo, tokenType string) error {
//反序列化标识
	identity, err := p.Deserializer.DeserializeIdentity(creator.Public())
	if err != nil {
		return errors.Wrapf(err, "identity [0x%x] cannot be deserialised", creator.Public())
	}

//检查身份有效性-在这个简单的策略中，所有有效的身份都是颁发者。
	if err := identity.Validate(); err != nil {
		return errors.Wrapf(err, "identity [0x%x] cannot be validated", creator.Public())
	}

	return nil
}

