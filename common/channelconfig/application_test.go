
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455945388298240>

/*
版权所有IBM Corp.2017保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channelconfig

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/capabilities"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/utils"
	. "github.com/onsi/gomega"
)

func TestApplicationInterface(t *testing.T) {
	_ = Application((*ApplicationConfig)(nil))
}

func TestACL(t *testing.T) {
	g := NewGomegaWithT(t)
	cgt := &cb.ConfigGroup{
		Values: map[string]*cb.ConfigValue{
			ACLsKey: {
				Value: utils.MarshalOrPanic(
					ACLValues(map[string]string{}).Value(),
				),
			},
			CapabilitiesKey: {
				Value: utils.MarshalOrPanic(
					CapabilitiesValue(map[string]bool{
						capabilities.ApplicationV1_2: true,
					}).Value(),
				),
			},
		},
	}

	t.Run("Success", func(t *testing.T) {
		cg := proto.Clone(cgt).(*cb.ConfigGroup)
		_, err := NewApplicationConfig(proto.Clone(cg).(*cb.ConfigGroup), nil)
		g.Expect(err).NotTo(HaveOccurred())
	})

	t.Run("MissingCapability", func(t *testing.T) {
		cg := proto.Clone(cgt).(*cb.ConfigGroup)
		delete(cg.Values, CapabilitiesKey)
		_, err := NewApplicationConfig(cg, nil)
		g.Expect(err).To(MatchError("ACLs may not be specified without the required capability"))
	})
}

