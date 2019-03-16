
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456077882167296>

/*
版权所有IBM公司保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package nwo

import (
	"fmt"
	"os"
	"runtime"

	"github.com/hyperledger/fabric/integration/helpers"
	"github.com/hyperledger/fabric/integration/runner"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

type Components struct {
	Paths map[string]string
}

var RequiredImages = []string{
	fmt.Sprintf("hyperledger/fabric-ccenv:%s-latest", runtime.GOARCH),
	runner.CouchDBDefaultImage,
	runner.KafkaDefaultImage,
	runner.ZooKeeperDefaultImage,
}

func (c *Components) Build(args ...string) {
	helpers.AssertImagesExist(RequiredImages...)

	if c.Paths == nil {
		c.Paths = map[string]string{}
	}
	cryptogen, err := gexec.Build("github.com/hyperledger/fabric/common/tools/cryptogen", args...)
	Expect(err).NotTo(HaveOccurred())
	c.Paths["cryptogen"] = cryptogen

	idemixgen, err := gexec.Build("github.com/hyperledger/fabric/common/tools/idemixgen", args...)
	Expect(err).NotTo(HaveOccurred())
	c.Paths["idemixgen"] = idemixgen

	configtxgen, err := gexec.Build("github.com/hyperledger/fabric/common/tools/configtxgen", args...)
	Expect(err).NotTo(HaveOccurred())
	c.Paths["configtxgen"] = configtxgen

	orderer, err := gexec.Build("github.com/hyperledger/fabric/orderer", args...)
	Expect(err).NotTo(HaveOccurred())
	c.Paths["orderer"] = orderer

	peer, err := gexec.Build("github.com/hyperledger/fabric/peer", args...)
	Expect(err).NotTo(HaveOccurred())
	c.Paths["peer"] = peer

	discover, err := gexec.Build("github.com/hyperledger/fabric/cmd/discover", args...)
	Expect(err).NotTo(HaveOccurred())
	c.Paths["discover"] = discover
}

func (c *Components) Cleanup() {
	for _, path := range c.Paths {
		err := os.Remove(path)
		Expect(err).NotTo(HaveOccurred())
	}
	gexec.CleanupBuildArtifacts()
}

func (c *Components) Cryptogen() string   { return c.Paths["cryptogen"] }
func (c *Components) Idemixgen() string   { return c.Paths["idemixgen"] }
func (c *Components) ConfigTxGen() string { return c.Paths["configtxgen"] }
func (c *Components) Orderer() string     { return c.Paths["orderer"] }
func (c *Components) Peer() string        { return c.Paths["peer"] }
func (c *Components) Discover() string    { return c.Paths["discover"] }

