
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:07</date>
//</624456000715362304>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package ccprovider

import (
	"bytes"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/common/sysccprovider"
)

//如果部署了具有给定名称和版本的链代码，is chaincode deployed将返回true。
func IsChaincodeDeployed(chainid, ccName, ccVersion string, ccHash []byte, sccp sysccprovider.SystemChaincodeProvider) (bool, error) {
	qe, err := sccp.GetQueryExecutorForLedger(chainid)
	if err != nil {
		return false, fmt.Errorf("Could not retrieve QueryExecutor for channel %s, error %s", chainid, err)
	}
	defer qe.Done()

//我们正在将LSCC表结构的细节泄漏到代码的其他部分，这很糟糕。
	chaincodeDataBytes, err := qe.GetState("lscc", ccName)
	if err != nil {
		return false, fmt.Errorf("Could not retrieve state for chaincode %s on channel %s, error %s", ccName, chainid, err)
	}

	if chaincodeDataBytes == nil {
		return false, nil
	}

	chaincodeData := &ChaincodeData{}
	err = proto.Unmarshal(chaincodeDataBytes, chaincodeData)
	if err != nil {
		return false, fmt.Errorf("Unmarshalling ChaincodeQueryResponse failed, error %s", err)
	}
	return chaincodeData.CCVersion() == ccVersion && bytes.Equal(chaincodeData.Hash(), ccHash), nil
}

