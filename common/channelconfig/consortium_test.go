
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455946428485632>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package channelconfig

import (
	"testing"

	"github.com/hyperledger/fabric/msp"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/stretchr/testify/assert"
)

func TestConsortiumConfig(t *testing.T) {
	cc, err := NewConsortiumConfig(&cb.ConfigGroup{}, NewMSPConfigHandler(msp.MSPv1_0))
	assert.NoError(t, err)
	orgs := cc.Organizations()
	assert.Equal(t, 0, len(orgs))

	policy := cc.ChannelCreationPolicy()
	assert.EqualValues(t, cb.Policy_UNKNOWN, policy.Type, "Expected policy type to be UNKNOWN")
}

