
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455947061825536>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channelconfig

import (
	"fmt"

	"github.com/hyperledger/fabric/msp"
	cb "github.com/hyperledger/fabric/protos/common"
	mspprotos "github.com/hyperledger/fabric/protos/msp"
	"github.com/pkg/errors"
)

const (
//msp key是排序器组中msp定义的键
	MSPKey = "MSP"
)

//OrganizationProtos用于反序列化组织配置
type OrganizationProtos struct {
	MSP *mspprotos.MSPConfig
}

//OrganizationConfig存储组织的配置
type OrganizationConfig struct {
	protos *OrganizationProtos

	mspConfigHandler *MSPConfigHandler
	msp              msp.MSP
	mspID            string
	name             string
}

//newOrganizationConfig为组织创建新的配置
func NewOrganizationConfig(name string, orgGroup *cb.ConfigGroup, mspConfigHandler *MSPConfigHandler) (*OrganizationConfig, error) {
	if len(orgGroup.Groups) > 0 {
		return nil, fmt.Errorf("organizations do not support sub-groups")
	}

	oc := &OrganizationConfig{
		protos:           &OrganizationProtos{},
		name:             name,
		mspConfigHandler: mspConfigHandler,
	}

	if err := DeserializeProtoValuesFromGroup(orgGroup, oc.protos); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize values")
	}

	if err := oc.Validate(); err != nil {
		return nil, err
	}

	return oc, nil
}

//name返回这个组织在config中引用的名称
func (oc *OrganizationConfig) Name() string {
	return oc.name
}

//msp id返回与此组织关联的msp id
func (oc *OrganizationConfig) MSPID() string {
	return oc.mspID
}

//validate返回配置是否有效
func (oc *OrganizationConfig) Validate() error {
	return oc.validateMSP()
}

func (oc *OrganizationConfig) validateMSP() error {
	var err error

	logger.Debugf("Setting up MSP for org %s", oc.name)
	oc.msp, err = oc.mspConfigHandler.ProposeMSP(oc.protos.MSP)
	if err != nil {
		return err
	}

	oc.mspID, _ = oc.msp.GetIdentifier()

	if oc.mspID == "" {
		return fmt.Errorf("MSP for org %s has empty MSP ID", oc.name)
	}

	return nil
}

