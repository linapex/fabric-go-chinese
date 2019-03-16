
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:26</date>
//</624456082680451072>

/*
版权所有IBM Corp.2017保留所有权利。

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


package mgmt

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/msp"
	"github.com/pkg/errors"
)

const (
//admins是本地MSP管理员的标签
	Admins = "Admins"

//成员是本地MSP成员的标签
	Members = "Members"
)

type MSPPrincipalGetter interface {
//get返回给定角色的MSP主体
	Get(role string) (*msp.MSPPrincipal, error)
}

func NewLocalMSPPrincipalGetter() MSPPrincipalGetter {
	return &localMSPPrincipalGetter{}
}

type localMSPPrincipalGetter struct{}

func (m *localMSPPrincipalGetter) Get(role string) (*msp.MSPPrincipal, error) {
	mspid, err := GetLocalMSP().GetIdentifier()
	if err != nil {
		return nil, errors.WithMessage(err, "could not extract local msp identifier")
	}

	switch role {
	case Admins:
		principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_ADMIN, MspIdentifier: mspid})
		if err != nil {
			return nil, errors.Wrap(err, "marshalling failed")
		}

		return &msp.MSPPrincipal{
			PrincipalClassification: msp.MSPPrincipal_ROLE,
			Principal:               principalBytes}, nil
	case Members:
		principalBytes, err := proto.Marshal(&msp.MSPRole{Role: msp.MSPRole_MEMBER, MspIdentifier: mspid})
		if err != nil {
			return nil, errors.Wrap(err, "marshalling failed")
		}

		return &msp.MSPPrincipal{
			PrincipalClassification: msp.MSPPrincipal_ROLE,
			Principal:               principalBytes}, nil
	default:
		return nil, errors.Errorf("MSP Principal role [%s] not recognized", role)
	}
}

