
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:29</date>
//</624456094386753536>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package msgprocessor

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/common/flogging"
	mockchannelconfig "github.com/hyperledger/fabric/common/mocks/config"
	mockpolicies "github.com/hyperledger/fabric/common/mocks/policies"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/utils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func init() {
	flogging.ActivateSpec("orderer.common.msgprocessor=DEBUG")
}

func makeEnvelope() *cb.Envelope {
	return &cb.Envelope{
		Payload: utils.MarshalOrPanic(&cb.Payload{
			Header: &cb.Header{
				SignatureHeader: utils.MarshalOrPanic(&cb.SignatureHeader{}),
			},
		}),
	}
}

func TestAccept(t *testing.T) {
	mpm := &mockchannelconfig.Resources{
		PolicyManagerVal: &mockpolicies.Manager{Policy: &mockpolicies.Policy{}},
	}
	assert.Nil(t, NewSigFilter("foo", mpm).Apply(makeEnvelope()), "Valid envelope and good policy")
}

func TestMissingPolicy(t *testing.T) {
	mpm := &mockchannelconfig.Resources{
		PolicyManagerVal: &mockpolicies.Manager{},
	}
	err := NewSigFilter("foo", mpm).Apply(makeEnvelope())
	assert.NotNil(t, err)
	assert.Regexp(t, "could not find policy", err.Error())
}

func TestEmptyPayload(t *testing.T) {
	mpm := &mockchannelconfig.Resources{
		PolicyManagerVal: &mockpolicies.Manager{Policy: &mockpolicies.Policy{}},
	}
	err := NewSigFilter("foo", mpm).Apply(&cb.Envelope{})
	assert.NotNil(t, err)
	assert.Regexp(t, "could not convert message to signedData", err.Error())
}

func TestErrorOnPolicy(t *testing.T) {
	mpm := &mockchannelconfig.Resources{
		PolicyManagerVal: &mockpolicies.Manager{Policy: &mockpolicies.Policy{Err: fmt.Errorf("Error")}},
	}
	err := NewSigFilter("foo", mpm).Apply(makeEnvelope())
	assert.NotNil(t, err)
	assert.Equal(t, ErrPermissionDenied, errors.Cause(err))
}

