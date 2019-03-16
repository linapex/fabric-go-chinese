
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455941940580352>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package main

import (
	"os/exec"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

func TestMissingArguments(t *testing.T) {
	gt := NewGomegaWithT(t)
	discover, err := Build("github.com/hyperledger/fabric/cmd/discover")
	gt.Expect(err).NotTo(HaveOccurred())
	defer CleanupBuildArtifacts()

//缺少密钥和证书标志
	cmd := exec.Command(discover, "--configFile", "conf.yaml", "--MSP", "SampleOrg", "saveConfig")
	process, err := Start(cmd, nil, nil)
	gt.Expect(err).NotTo(HaveOccurred())
	gt.Eventually(process).Should(Exit(1))
	gt.Expect(process.Err).To(gbytes.Say("empty string that is mandatory"))
}

