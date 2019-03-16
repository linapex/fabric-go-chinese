
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455967995596800>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package policies

import (
	"testing"

	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/stretchr/testify/assert"
)

func basicTest(t *testing.T, sv *StandardConfigPolicy) {
	assert.NotNil(t, sv)
	assert.NotEmpty(t, sv.Key())
	assert.NotNil(t, sv.Value())
}

func TestUtilsBasic(t *testing.T) {
	basicTest(t, ImplicitMetaAnyPolicy("foo"))
	basicTest(t, ImplicitMetaAllPolicy("foo"))
	basicTest(t, ImplicitMetaMajorityPolicy("foo"))
	basicTest(t, SignaturePolicy("foo", &cb.SignaturePolicyEnvelope{}))
}

