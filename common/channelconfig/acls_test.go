
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455944838844416>

/*
版权所有State Street Corp.2018保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channelconfig

import (
	"testing"

	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
)

const (
	sampleAPI1Name      = "Foo"
	sampleAPI1PolicyRef = "foo"

	sampleAPI2Name      = "Bar"
	sampleAPI2PolicyRef = "/Channel/foo"
)

var sampleAPIsProvider = map[string]*pb.APIResource{
	sampleAPI1Name: {PolicyRef: sampleAPI1PolicyRef},
	sampleAPI2Name: {PolicyRef: sampleAPI2PolicyRef},
}

func TestGreenAPIsPath(t *testing.T) {
	ag := newAPIsProvider(sampleAPIsProvider)
	assert.NotNil(t, ag)

	t.Run("PresentAPIs", func(t *testing.T) {
		assert.Equal(t, "/Channel/Application/"+sampleAPI1PolicyRef, ag.PolicyRefForAPI(sampleAPI1Name))
		assert.Equal(t, sampleAPI2PolicyRef, ag.PolicyRefForAPI(sampleAPI2Name))
	})

	t.Run("MissingAPIs", func(t *testing.T) {
		assert.Empty(t, ag.PolicyRefForAPI("missing"))
	})
}

func TestNilACLs(t *testing.T) {
	ccg := newAPIsProvider(nil)

	assert.NotNil(t, ccg)
	assert.NotNil(t, ccg.aclPolicyRefs)
	assert.Empty(t, ccg.aclPolicyRefs)
}

func TestEmptyACLs(t *testing.T) {
	ccg := newAPIsProvider(map[string]*pb.APIResource{})

	assert.NotNil(t, ccg)
	assert.NotNil(t, ccg.aclPolicyRefs)
	assert.Empty(t, ccg.aclPolicyRefs)
}

