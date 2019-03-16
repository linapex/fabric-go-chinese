
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:29</date>
//</624456094156066816>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package msgprocessor

import (
	"fmt"
	"testing"

	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/stretchr/testify/assert"
)

var RejectRule = Rule(rejectRule{})

type rejectRule struct{}

func (r rejectRule) Apply(message *cb.Envelope) error {
	return fmt.Errorf("Rejected")
}

func TestEmptyRejectRule(t *testing.T) {
	t.Run("Reject", func(t *testing.T) {
		assert.NotNil(t, EmptyRejectRule.Apply(&cb.Envelope{}))
	})
	t.Run("Accept", func(t *testing.T) {
		assert.Nil(t, EmptyRejectRule.Apply(&cb.Envelope{Payload: []byte("fakedata")}))
	})
}

func TestAcceptRule(t *testing.T) {
	assert.Nil(t, AcceptRule.Apply(&cb.Envelope{}))
}

func TestRuleSet(t *testing.T) {
	t.Run("RejectAccept", func(t *testing.T) {
		assert.NotNil(t, NewRuleSet([]Rule{RejectRule, AcceptRule}).Apply(&cb.Envelope{}))
	})
	t.Run("AcceptReject", func(t *testing.T) {
		assert.NotNil(t, NewRuleSet([]Rule{AcceptRule, RejectRule}).Apply(&cb.Envelope{}))
	})
	t.Run("AcceptAccept", func(t *testing.T) {
		assert.Nil(t, NewRuleSet([]Rule{AcceptRule, AcceptRule}).Apply(&cb.Envelope{}))
	})
	t.Run("Empty", func(t *testing.T) {
		assert.Nil(t, NewRuleSet(nil).Apply(&cb.Envelope{}))
	})
}

