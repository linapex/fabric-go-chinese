
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:36</date>
//</624456123059015680>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package utils

import (
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
)

//取消标记haincodedeploymentspec取消标记chaincodedeploymentspec自
//提供的字节
func UnmarshalChaincodeDeploymentSpec(cdsBytes []byte) (*peer.ChaincodeDeploymentSpec, error) {
	cds := &peer.ChaincodeDeploymentSpec{}
	err := proto.Unmarshal(cdsBytes, cds)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshaling ChaincodeDeploymentSpec")
	}

	return cds, nil
}

