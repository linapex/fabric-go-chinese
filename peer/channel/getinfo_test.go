
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:32</date>
//</624456108408311808>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channel

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/peer/common"
	cb "github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
)

func TestGetChannelInfo(t *testing.T) {
	InitMSP()
	resetFlags()

	mockBlockchainInfo := &cb.BlockchainInfo{
		Height:            1,
		CurrentBlockHash:  []byte("CurrentBlockHash"),
		PreviousBlockHash: []byte("PreviousBlockHash"),
	}
	mockPayload, err := proto.Marshal(mockBlockchainInfo)
	assert.NoError(t, err)

	mockResponse := &pb.ProposalResponse{
		Response: &pb.Response{
			Status:  200,
			Payload: mockPayload,
		},
		Endorsement: &pb.Endorsement{},
	}

	signer, err := common.GetDefaultSigner()
	assert.NoError(t, err)

	mockCF := &ChannelCmdFactory{
		EndorserClient:   common.GetMockEndorserClient(mockResponse, nil),
		BroadcastFactory: mockBroadcastClientFactory,
		Signer:           signer,
	}

	cmd := getinfoCmd(mockCF)
	AddFlags(cmd)

	args := []string{"-c", mockChannel}
	cmd.SetArgs(args)

	assert.NoError(t, cmd.Execute())
}

func TestGetChannelInfoMissingChannelID(t *testing.T) {
	InitMSP()
	resetFlags()

	signer, err := common.GetDefaultSigner()
	if err != nil {
		t.Fatalf("Get default signer error: %v", err)
	}

	mockCF := &ChannelCmdFactory{
		Signer: signer,
	}

	cmd := getinfoCmd(mockCF)

	AddFlags(cmd)

	cmd.SetArgs([]string{})

	assert.Error(t, cmd.Execute())
}

