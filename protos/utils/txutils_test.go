
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:36</date>
//</624456123654606848>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package utils_test

import (
	"encoding/hex"
	"errors"
	"strconv"
	"testing"

	"github.com/golang/protobuf/proto"
	mockmsp "github.com/hyperledger/fabric/common/mocks/msp"
	"github.com/hyperledger/fabric/common/util"
	cb "github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetPayloads(t *testing.T) {
	var txAction *pb.TransactionAction
	var err error

//好的
	ccActionBytes, _ := proto.Marshal(&pb.ChaincodeAction{
		Results: []byte("results"),
	})
	proposalResponsePayload := &pb.ProposalResponsePayload{
		Extension: ccActionBytes,
	}
	proposalResponseBytes, err := proto.Marshal(proposalResponsePayload)
	ccActionPayload := &pb.ChaincodeActionPayload{
		Action: &pb.ChaincodeEndorsedAction{
			ProposalResponsePayload: proposalResponseBytes,
		},
	}
	ccActionPayloadBytes, _ := proto.Marshal(ccActionPayload)
	txAction = &pb.TransactionAction{
		Payload: ccActionPayloadBytes,
	}
	_, _, err = utils.GetPayloads(txAction)
	assert.NoError(t, err, "Unexpected error getting payload bytes")
	t.Logf("error1 [%s]", err)

//无提案回复延期
	proposalResponseBytes, err = proto.Marshal(&pb.ProposalResponsePayload{
		Extension: nil,
	})
	ccActionPayloadBytes, _ = proto.Marshal(&pb.ChaincodeActionPayload{
		Action: &pb.ChaincodeEndorsedAction{
			ProposalResponsePayload: proposalResponseBytes,
		},
	})
	txAction = &pb.TransactionAction{
		Payload: ccActionPayloadBytes,
	}
	_, _, err = utils.GetPayloads(txAction)
	assert.Error(t, err, "Expected error with nil proposal response extension")
	t.Logf("error2 [%s]", err)

//格式错误的建议响应负载
	ccActionPayloadBytes, _ = proto.Marshal(&pb.ChaincodeActionPayload{
		Action: &pb.ChaincodeEndorsedAction{
			ProposalResponsePayload: []byte("bad payload"),
		},
	})
	txAction = &pb.TransactionAction{
		Payload: ccActionPayloadBytes,
	}
	_, _, err = utils.GetPayloads(txAction)
	assert.Error(t, err, "Expected error with malformed proposal response payload")
	t.Logf("error3 [%s]", err)

//格式错误的建议响应负载扩展
	proposalResponseBytes, _ = proto.Marshal(&pb.ProposalResponsePayload{
		Extension: []byte("bad extension"),
	})
	ccActionPayloadBytes, _ = proto.Marshal(&pb.ChaincodeActionPayload{
		Action: &pb.ChaincodeEndorsedAction{
			ProposalResponsePayload: proposalResponseBytes,
		},
	})
	txAction = &pb.TransactionAction{
		Payload: ccActionPayloadBytes,
	}
	_, _, err = utils.GetPayloads(txAction)
	assert.Error(t, err, "Expected error with malformed proposal response extension")
	t.Logf("error4 [%s]", err)

//无提案响应有效载荷扩展
	proposalResponseBytes, _ = proto.Marshal(&pb.ProposalResponsePayload{
		ProposalHash: []byte("hash"),
	})
	ccActionPayloadBytes, _ = proto.Marshal(&pb.ChaincodeActionPayload{
		Action: &pb.ChaincodeEndorsedAction{
			ProposalResponsePayload: proposalResponseBytes,
		},
	})
	txAction = &pb.TransactionAction{
		Payload: ccActionPayloadBytes,
	}
	_, _, err = utils.GetPayloads(txAction)
	assert.Error(t, err, "Expected error with nil proposal response extension")
	t.Logf("error5 [%s]", err)

//错误的事务操作负载
	txAction = &pb.TransactionAction{
		Payload: []byte("bad payload"),
	}
	_, _, err = utils.GetPayloads(txAction)
	assert.Error(t, err, "Expected error with malformed transaction action payload")
	t.Logf("error6 [%s]", err)

}

func TestCreateSignedTx(t *testing.T) {
	var err error
	prop := &pb.Proposal{}

	signID, err := mockmsp.NewNoopMsp().GetDefaultSigningIdentity()
	assert.NoError(t, err, "Unexpected error getting signing identity")
	signerBytes, err := signID.Serialize()
	assert.NoError(t, err, "Unexpected error serializing signing identity")

	ccHeaderExtensionBytes, _ := proto.Marshal(&pb.ChaincodeHeaderExtension{})
	chdrBytes, _ := proto.Marshal(&cb.ChannelHeader{
		Extension: ccHeaderExtensionBytes,
	})
	shdrBytes, _ := proto.Marshal(&cb.SignatureHeader{
		Creator: signerBytes,
	})
	responses := []*pb.ProposalResponse{{}}

//链码头扩展格式不正确
	headerBytes, _ := proto.Marshal(&cb.Header{
		ChannelHeader:   []byte("bad channel header"),
		SignatureHeader: shdrBytes,
	})
	prop.Header = headerBytes
	_, err = utils.CreateSignedTx(prop, signID, responses...)
	assert.Error(t, err, "Expected error with malformed chaincode extension")

//签名头格式不正确
	headerBytes, _ = proto.Marshal(&cb.Header{
		SignatureHeader: []byte("bad signature header"),
	})
	prop.Header = headerBytes
	_, err = utils.CreateSignedTx(prop, signID, responses...)
	assert.Error(t, err, "Expected error with malformed signature header")

//设置其余测试的头字节
	headerBytes, _ = proto.Marshal(&cb.Header{
		ChannelHeader:   chdrBytes,
		SignatureHeader: shdrBytes,
	})
	prop.Header = headerBytes

//不匹配的响应
	responses = []*pb.ProposalResponse{{
		Payload: []byte("payload"),
		Response: &pb.Response{
			Status: int32(200),
		},
	}}
	responses = append(responses, &pb.ProposalResponse{
		Payload: []byte("payload2"),
		Response: &pb.Response{
			Status: int32(200),
		},
	})
	_, err = utils.CreateSignedTx(prop, signID, responses...)
	assert.Error(t, err, "Expected error with non-matching responses")

//无背书
	responses = []*pb.ProposalResponse{{
		Payload: []byte("payload"),
		Response: &pb.Response{
			Status: int32(200),
		},
	}}
	_, err = utils.CreateSignedTx(prop, signID, responses...)
	assert.Error(t, err, "Expected error with no endorsements")

//成功
	responses = []*pb.ProposalResponse{{
		Payload:     []byte("payload"),
		Endorsement: &pb.Endorsement{},
		Response: &pb.Response{
			Status: int32(200),
		},
	}}
	_, err = utils.CreateSignedTx(prop, signID, responses...)
	assert.NoError(t, err, "Unexpected error creating signed transaction")
	t.Logf("error: [%s]", err)

//
//
//其他故障案例
	prop = &pb.Proposal{}
	responses = []*pb.ProposalResponse{}
//无提案响应
	_, err = utils.CreateSignedTx(prop, signID, responses...)
	assert.Error(t, err, "Expected error with no proposal responses")

//缺少建议标题
	responses = append(responses, &pb.ProposalResponse{})
	_, err = utils.CreateSignedTx(prop, signID, responses...)
	assert.Error(t, err, "Expected error with no proposal header")

//错误的建议负载
	prop.Payload = []byte("bad payload")
	_, err = utils.CreateSignedTx(prop, signID, responses...)
	assert.Error(t, err, "Expected error with malformed proposal payload")

//无效负载头
	prop.Header = []byte("bad header")
	_, err = utils.CreateSignedTx(prop, signID, responses...)
	assert.Error(t, err, "Expected error with malformed proposal header")

}

func TestCreateSignedTxStatus(t *testing.T) {
	serializedExtension, err := proto.Marshal(&pb.ChaincodeHeaderExtension{})
	assert.NoError(t, err)
	serializedChannelHeader, err := proto.Marshal(&cb.ChannelHeader{
		Extension: serializedExtension,
	})
	assert.NoError(t, err)

	signingID, err := mockmsp.NewNoopMsp().GetDefaultSigningIdentity()
	assert.NoError(t, err)
	serializedSigningID, err := signingID.Serialize()
	assert.NoError(t, err)
	serializedSignatureHeader, err := proto.Marshal(&cb.SignatureHeader{
		Creator: serializedSigningID,
	})
	assert.NoError(t, err)

	header := &cb.Header{
		ChannelHeader:   serializedChannelHeader,
		SignatureHeader: serializedSignatureHeader,
	}

	serializedHeader, err := proto.Marshal(header)
	assert.NoError(t, err)

	proposal := &pb.Proposal{
		Header: serializedHeader,
	}

	tests := []struct {
		status      int32
		expectedErr string
	}{
		{status: 0, expectedErr: "proposal response was not successful, error code 0, msg response-message"},
		{status: 199, expectedErr: "proposal response was not successful, error code 199, msg response-message"},
		{status: 200, expectedErr: ""},
		{status: 201, expectedErr: ""},
		{status: 399, expectedErr: ""},
		{status: 400, expectedErr: "proposal response was not successful, error code 400, msg response-message"},
	}
	for _, tc := range tests {
		t.Run(strconv.Itoa(int(tc.status)), func(t *testing.T) {
			response := &pb.ProposalResponse{
				Payload:     []byte("payload"),
				Endorsement: &pb.Endorsement{},
				Response: &pb.Response{
					Status:  tc.status,
					Message: "response-message",
				},
			}

			_, err := utils.CreateSignedTx(proposal, signingID, response)
			if tc.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr)
			}
		})
	}
}

func TestCreateSignedEnvelope(t *testing.T) {
	var env *cb.Envelope
	channelID := "mychannelID"
	msg := &cb.ConfigEnvelope{}

	env, err := utils.CreateSignedEnvelope(cb.HeaderType_CONFIG, channelID,
		goodSigner, msg, int32(1), uint64(1))
	assert.NoError(t, err, "Unexpected error creating signed envelope")
	assert.NotNil(t, env, "Envelope should not be nil")
//模拟符号返回要签名的字节
	assert.Equal(t, env.Payload, env.Signature, "Unexpected signature returned")
	payload := &cb.Payload{}
	err = proto.Unmarshal(env.Payload, payload)
	assert.NoError(t, err, "Failed to unmarshal payload")
	data := &cb.ConfigEnvelope{}
	err = proto.Unmarshal(payload.Data, data)
	assert.NoError(t, err, "Expected payload data to be a config envelope")
	assert.Equal(t, msg, data, "Payload data does not match expected value")

	_, err = utils.CreateSignedEnvelope(cb.HeaderType_CONFIG, channelID,
		badSigner, &cb.ConfigEnvelope{}, int32(1), uint64(1))
	assert.Error(t, err, "Expected sign error")
}

func TestCreateSignedEnvelopeNilSigner(t *testing.T) {
	var env *cb.Envelope
	channelID := "mychannelID"
	msg := &cb.ConfigEnvelope{}

	env, err := utils.CreateSignedEnvelope(cb.HeaderType_CONFIG, channelID,
		nil, msg, int32(1), uint64(1))
	assert.NoError(t, err, "Unexpected error creating signed envelope")
	assert.NotNil(t, env, "Envelope should not be nil")
	assert.Empty(t, env.Signature, "Signature should have been empty")
	payload := &cb.Payload{}
	err = proto.Unmarshal(env.Payload, payload)
	assert.NoError(t, err, "Failed to unmarshal payload")
	data := &cb.ConfigEnvelope{}
	err = proto.Unmarshal(payload.Data, data)
	assert.NoError(t, err, "Expected payload data to be a config envelope")
	assert.Equal(t, msg, data, "Payload data does not match expected value")
}

func TestGetSignedProposal(t *testing.T) {
	var signedProp *pb.SignedProposal
	var err error

	signID, err := mockmsp.NewNoopMsp().GetDefaultSigningIdentity()
	assert.NoError(t, err, "Unexpected error getting signing identity")

	prop := &pb.Proposal{}
	propBytes, _ := proto.Marshal(prop)
	signedProp, err = utils.GetSignedProposal(prop, signID)
	assert.NoError(t, err, "Unexpected error getting signed proposal")
	assert.Equal(t, propBytes, signedProp.ProposalBytes,
		"Proposal bytes did not match expected value")
	assert.Equal(t, []byte("signature"), signedProp.Signature,
		"Signature did not match expected value")

	_, err = utils.GetSignedProposal(nil, signID)
	assert.Error(t, err, "Expected error with nil proposal")
	_, err = utils.GetSignedProposal(prop, nil)
	assert.Error(t, err, "Expected error with nil signing identity")

}

func TestMockSignedEndorserProposalOrPanic(t *testing.T) {
	var prop *pb.Proposal
	var signedProp *pb.SignedProposal

	ccProposal := &pb.ChaincodeProposalPayload{}
	cis := &pb.ChaincodeInvocationSpec{}
	chainID := "testchainid"
	sig := []byte("signature")
	creator := []byte("creator")
	cs := &pb.ChaincodeSpec{
		ChaincodeId: &pb.ChaincodeID{
			Name: "mychaincode",
		},
	}

	signedProp, prop = utils.MockSignedEndorserProposalOrPanic(chainID, cs,
		creator, sig)
	assert.Equal(t, sig, signedProp.Signature,
		"Signature did not match expected result")
	propBytes, _ := proto.Marshal(prop)
	assert.Equal(t, propBytes, signedProp.ProposalBytes,
		"Proposal bytes do not match expected value")
	err := proto.Unmarshal(prop.Payload, ccProposal)
	assert.NoError(t, err, "Expected ChaincodeProposalPayload")
	err = proto.Unmarshal(ccProposal.Input, cis)
	assert.NoError(t, err, "Expected ChaincodeInvocationSpec")
	assert.Equal(t, cs.ChaincodeId.Name, cis.ChaincodeSpec.ChaincodeId.Name,
		"Chaincode name did not match expected value")
}

func TestMockSignedEndorserProposal2OrPanic(t *testing.T) {
	var prop *pb.Proposal
	var signedProp *pb.SignedProposal

	ccProposal := &pb.ChaincodeProposalPayload{}
	cis := &pb.ChaincodeInvocationSpec{}
	chainID := "testchainid"
	sig := []byte("signature")
	signID, err := mockmsp.NewNoopMsp().GetDefaultSigningIdentity()
	assert.NoError(t, err, "Unexpected error getting signing identity")

	signedProp, prop = utils.MockSignedEndorserProposal2OrPanic(chainID,
		&pb.ChaincodeSpec{}, signID)
	assert.Equal(t, sig, signedProp.Signature,
		"Signature did not match expected result")
	propBytes, _ := proto.Marshal(prop)
	assert.Equal(t, propBytes, signedProp.ProposalBytes,
		"Proposal bytes do not match expected value")
	err = proto.Unmarshal(prop.Payload, ccProposal)
	assert.NoError(t, err, "Expected ChaincodeProposalPayload")
	err = proto.Unmarshal(ccProposal.Input, cis)
	assert.NoError(t, err, "Expected ChaincodeInvocationSpec")
}

func TestGetBytesProposalPayloadForTx(t *testing.T) {
	input := &pb.ChaincodeProposalPayload{
		Input:        []byte("input"),
		TransientMap: make(map[string][]byte),
	}
	expected, _ := proto.Marshal(&pb.ChaincodeProposalPayload{
		Input: []byte("input"),
	})

	result, err := utils.GetBytesProposalPayloadForTx(input, []byte{})
	assert.NoError(t, err, "Unexpected error getting proposal payload")
	assert.Equal(t, expected, result, "Payload does not match expected value")

	_, err = utils.GetBytesProposalPayloadForTx(nil, []byte{})
	assert.Error(t, err, "Expected error with nil proposal payload")
}

func TestGetProposalHash2(t *testing.T) {
	expectedHashHex := "7b622ef4e1ab9b7093ec3bbfbca17d5d6f14a437914a6839319978a7034f7960"
	expectedHash, _ := hex.DecodeString(expectedHashHex)
	hdr := &cb.Header{
		ChannelHeader:   []byte("chdr"),
		SignatureHeader: []byte("shdr"),
	}
	propHash, err := utils.GetProposalHash2(hdr, []byte("ccproppayload"))
	assert.NoError(t, err, "Unexpected error getting hash2 for proposal")
	t.Logf("%x", propHash)
	assert.Equal(t, expectedHash, propHash,
		"Proposal hash did not match expected hash")

	propHash, err = utils.GetProposalHash2(&cb.Header{},
		[]byte("ccproppayload"))
	assert.Error(t, err, "Expected error with nil arguments")
}

func TestGetProposalHash1(t *testing.T) {
	expectedHashHex := "d4c1e3cac2105da5fddc2cfe776d6ec28e4598cf1e6fa51122c7f70d8076437b"
	expectedHash, _ := hex.DecodeString(expectedHashHex)
	hdr := &cb.Header{
		ChannelHeader:   []byte("chdr"),
		SignatureHeader: []byte("shdr"),
	}

	ccProposal, _ := proto.Marshal(&pb.ChaincodeProposalPayload{})

	propHash, err := utils.GetProposalHash1(hdr, ccProposal, []byte{})
	assert.NoError(t, err, "Unexpected error getting hash for proposal")
	t.Logf("%x", propHash)
	assert.Equal(t, expectedHash, propHash,
		"Proposal hash did not match expected hash")

	propHash, err = utils.GetProposalHash1(hdr,
		[]byte("ccproppayload"), []byte{})
	assert.Error(t, err,
		"Expected error with malformed chaincode proposal payload")

	propHash, err = utils.GetProposalHash1(&cb.Header{},
		[]byte("ccproppayload"), []byte{})
	assert.Error(t, err, "Expected error with nil arguments")
}

func TestCreateProposalResponseFailure(t *testing.T) {
//从chaincodeinvocationspec创建建议
	prop, _, err := utils.CreateChaincodeProposal(cb.HeaderType_ENDORSER_TRANSACTION, util.GetTestChainID(), createCIS(), signerSerialized)
	if err != nil {
		t.Fatalf("Could not create chaincode proposal, err %s\n", err)
		return
	}

	response := &pb.Response{Status: 502, Payload: []byte("Invalid function name")}
	result := []byte("res")
	ccid := &pb.ChaincodeID{Name: "foo", Version: "v1"}

	prespFailure, err := utils.CreateProposalResponseFailure(prop.Header, prop.Payload, response, result, nil, ccid, nil)
	if err != nil {
		t.Fatalf("Could not create proposal response failure, err %s\n", err)
		return
	}

	assert.Equal(t, int32(502), prespFailure.Response.Status)
//深入到响应中以查找链码响应
	pRespPayload, err := utils.GetProposalResponsePayload(prespFailure.Payload)
	assert.NoError(t, err, "Error while unmarshaling proposal response payload: %s", err)
	ca, err := utils.GetChaincodeAction(pRespPayload.Extension)
	assert.NoError(t, err, "Error while unmarshaling chaincode action: %s", err)

	assert.Equal(t, int32(502), ca.Response.Status)
	assert.Equal(t, "Invalid function name", string(ca.Response.Payload))
}

//嘲弄
var badSigner = &mockLocalSigner{
	returnError: true,
}

var goodSigner = &mockLocalSigner{
	returnError: false,
}

type mockLocalSigner struct {
	returnError bool
}

func (m *mockLocalSigner) NewSignatureHeader() (*cb.SignatureHeader, error) {
	if m.returnError {
		return nil, errors.New("signature header error")
	}
	return &cb.SignatureHeader{}, nil
}

func (m *mockLocalSigner) Sign(message []byte) ([]byte, error) {
	if m.returnError {
		return nil, errors.New("sign error")
	}
	return message, nil
}

