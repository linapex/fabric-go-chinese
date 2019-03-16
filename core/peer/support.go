
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:16</date>
//</624456039571394560>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package peer

import (
	"github.com/hyperledger/fabric/common/channelconfig"
)

var supportFactory SupportFactory

//SupportFactory是支持接口的工厂
type SupportFactory interface {
//NewSupport返回支持接口
	NewSupport() Support
}

//支持允许访问对等资源并避免调用静态方法
type Support interface {
//GETAPPLICATIOFIGG返回通道的CONTXXAPPLATION.SARDCONFIGG
//以及应用程序配置是否存在
	GetApplicationConfig(cid string) (channelconfig.Application, bool)
}

type supportImpl struct {
	operations Operations
}

func (s *supportImpl) GetApplicationConfig(cid string) (channelconfig.Application, bool) {
	cc := s.operations.GetChannelConfig(cid)
	if cc == nil {
		return nil, false
	}

	return cc.ApplicationConfig()
}

