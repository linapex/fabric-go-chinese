
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:32</date>
//</624456106546040832>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package chaincode

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/peer/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
)

func TestChaincodeListCmd(t *testing.T) {
	signer, err := common.GetDefaultSigner()
	if err != nil {
		t.Fatalf("Get default signer error: %s", err)
	}

	installedCqr := &pb.ChaincodeQueryResponse{
		Chaincodes: []*pb.ChaincodeInfo{
			{Name: "mycc1", Version: "1.0", Path: "codePath1", Input: "input", Escc: "escc", Vscc: "vscc", Id: []byte{1, 2, 3}},
			{Name: "mycc2", Version: "1.0", Path: "codePath2", Input: "input", Escc: "escc", Vscc: "vscc"},
		},
	}
	installedCqrBytes, err := proto.Marshal(installedCqr)
	if err != nil {
		t.Fatalf("Marshale error: %s", err)
	}

	mockResponse := &pb.ProposalResponse{
		Response:    &pb.Response{Status: 200, Payload: installedCqrBytes},
		Endorsement: &pb.Endorsement{},
	}
	mockEndorserClients := []pb.EndorserClient{common.GetMockEndorserClient(mockResponse, nil)}
	mockBroadcastClient := common.GetMockBroadcastClient(nil)
	mockCF := &ChaincodeCmdFactory{
		EndorserClients: mockEndorserClients,
		Signer:          signer,
		BroadcastClient: mockBroadcastClient,
	}

//重置channelid，它可能是由以前的测试设置的
	channelID = ""

//获取已安装的链代码
	installedChaincodesCmd := listCmd(mockCF)

	args := []string{"--installed"}
	installedChaincodesCmd.SetArgs(args)
	if err := installedChaincodesCmd.Execute(); err != nil {
		t.Errorf("Run chaincode list cmd to get installed chaincodes error:%v", err)
	}

	resetFlags()

//获取实例化的链代码
	instantiatedChaincodesCmd := listCmd(mockCF)
	args = []string{"--instantiated"}
	instantiatedChaincodesCmd.SetArgs(args)
	err = instantiatedChaincodesCmd.Execute()
	assert.Error(t, err, "Run chaincode list cmd to get instantiated chaincodes should fail if invoked without -C flag")

	args = []string{"--instantiated", "-C", "mychannel"}
	instantiatedChaincodesCmd.SetArgs(args)
	if err := instantiatedChaincodesCmd.Execute(); err != nil {
		t.Errorf("Run chaincode list cmd to get instantiated chaincodes error:%v", err)
	}

	resetFlags()

//错误的情况：同时设置“-installed”和“-installed”
	Cmd := listCmd(mockCF)
	args = []string{"--installed", "--instantiated"}
	Cmd.SetArgs(args)
	err = Cmd.Execute()
	assert.Error(t, err, "Run chaincode list cmd to get instantiated/installed chaincodes should fail if invoked without -C flag")

	args = []string{"--installed", "--instantiated", "-C", "mychannel"}
	Cmd.SetArgs(args)
	expectErr := fmt.Errorf("Must explicitly specify \"--installed\" or \"--instantiated\"")
	if err := Cmd.Execute(); err == nil || err.Error() != expectErr.Error() {
		t.Errorf("Expect error: %s", expectErr)
	}

	resetFlags()

//错误的情况：Miss“—intsalled”和“—instanced”
	nilCmd := listCmd(mockCF)

	args = []string{"-C", "mychannel"}
	nilCmd.SetArgs(args)

	expectErr = fmt.Errorf("Must explicitly specify \"--installed\" or \"--instantiated\"")
	if err := nilCmd.Execute(); err == nil || err.Error() != expectErr.Error() {
		t.Errorf("Expect error: %s", expectErr)
	}
}

func TestChaincodeListFailure(t *testing.T) {
	signer, err := common.GetDefaultSigner()
	if err != nil {
		t.Fatalf("Get default signer error: %s", err)
	}

	mockResponse := &pb.ProposalResponse{
		Response:    &pb.Response{Status: 500, Message: "error message"},
		Endorsement: &pb.Endorsement{},
	}
	mockEndorserClients := []pb.EndorserClient{common.GetMockEndorserClient(mockResponse, nil)}
	mockBroadcastClient := common.GetMockBroadcastClient(nil)
	mockCF := &ChaincodeCmdFactory{
		EndorserClients: mockEndorserClients,
		Signer:          signer,
		BroadcastClient: mockBroadcastClient,
	}

//重置channelid，它可能是由以前的测试设置的
	channelID = ""

	resetFlags()

//获取实例化的链代码
	instantiatedChaincodesCmd := listCmd(mockCF)
	args := []string{"--instantiated", "-C", "mychannel"}
	instantiatedChaincodesCmd.SetArgs(args)
	err = instantiatedChaincodesCmd.Execute()
	assert.Error(t, err)
	assert.Regexp(t, "Bad response: 500 - error message", err.Error())
}

func TestString(t *testing.T) {
	id := []byte{1, 2, 3, 4, 5}
	idBytes := hex.EncodeToString(id)
	b, _ := hex.DecodeString(idBytes)
	ccInf := &ccInfo{
		ChaincodeInfo: &pb.ChaincodeInfo{
			Name:    "ccName",
			Id:      b,
			Version: "1.0",
			Escc:    "escc",
			Input:   "input",
			Vscc:    "vscc",
		},
	}
	assert.Equal(t, "Name: ccName, Version: 1.0, Input: input, Escc: escc, Vscc: vscc, Id: 0102030405", ccInf.String())
}

