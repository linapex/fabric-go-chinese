
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455945220526080>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channelconfig

import (
	"fmt"

	cb "github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

const (
//anchorpeersky是anchorpeers配置值的键名
	AnchorPeersKey = "AnchorPeers"
)

//应用程序或协议从配置反序列化
type ApplicationOrgProtos struct {
	AnchorPeers *pb.AnchorPeers
}

//applicationOrgConfig定义应用程序组织的配置
type ApplicationOrgConfig struct {
	*OrganizationConfig
	protos *ApplicationOrgProtos
	name   string
}

//NewApplicationOrgConfig为应用程序组织创建新的配置
func NewApplicationOrgConfig(id string, orgGroup *cb.ConfigGroup, mspConfig *MSPConfigHandler) (*ApplicationOrgConfig, error) {
	if len(orgGroup.Groups) > 0 {
		return nil, fmt.Errorf("ApplicationOrg config does not allow sub-groups")
	}

	protos := &ApplicationOrgProtos{}
	orgProtos := &OrganizationProtos{}

	if err := DeserializeProtoValuesFromGroup(orgGroup, protos, orgProtos); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize values")
	}

	aoc := &ApplicationOrgConfig{
		name:   id,
		protos: protos,
		OrganizationConfig: &OrganizationConfig{
			name:             id,
			protos:           orgProtos,
			mspConfigHandler: mspConfig,
		},
	}

	if err := aoc.Validate(); err != nil {
		return nil, err
	}

	return aoc, nil
}

//anchor peers返回此组织的锚定对等方列表
func (aog *ApplicationOrgConfig) AnchorPeers() []*pb.AnchorPeer {
	return aog.protos.AnchorPeers.AnchorPeers
}

func (aoc *ApplicationOrgConfig) Validate() error {
	logger.Debugf("Anchor peers for org %s are %v", aoc.name, aoc.protos.AnchorPeers)
	return aoc.OrganizationConfig.Validate()
}

