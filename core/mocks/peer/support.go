
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456037184835584>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package peer

import (
	"github.com/hyperledger/fabric/common/channelconfig"
	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/core/peer"
)

type MockSupportImpl struct {
	GetApplicationConfigRv     channelconfig.Application
	GetApplicationConfigBoolRv bool
	ChaincodeByNameRv          *ccprovider.ChaincodeData
	ChaincodeByNameBoolRv      bool
}

func (s *MockSupportImpl) GetApplicationConfig(cid string) (channelconfig.Application, bool) {
	return s.GetApplicationConfigRv, s.GetApplicationConfigBoolRv
}

func (s *MockSupportImpl) ChaincodeByName(chainname, ccname string) (*ccprovider.ChaincodeData, bool) {
	return s.ChaincodeByNameRv, s.ChaincodeByNameBoolRv
}

type MockSupportFactoryImpl struct {
	NewSupportRv *MockSupportImpl
}

func (c *MockSupportFactoryImpl) NewSupport() peer.Support {
	return c.NewSupportRv
}

