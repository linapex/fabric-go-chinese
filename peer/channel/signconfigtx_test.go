
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:32</date>
//</624456108957765632>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channel

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/hyperledger/fabric/peer/common"
	cb "github.com/hyperledger/fabric/protos/common"
	"github.com/stretchr/testify/assert"
)

func TestSignConfigtx(t *testing.T) {
	InitMSP()
	resetFlags()

	dir, err := ioutil.TempDir("/tmp", "signconfigtxtest-")
	if err != nil {
		t.Fatalf("couldn't create temp dir")
	}
defer os.RemoveAll(dir) //清理

	configtxFile := filepath.Join(dir, mockChannel)
	if _, err = createTxFile(configtxFile, cb.HeaderType_CONFIG_UPDATE, mockChannel); err != nil {
		t.Fatalf("couldn't create tx file")
	}

	signer, err := common.GetDefaultSigner()
	if err != nil {
		t.Fatalf("Get default signer error: %v", err)
	}

	mockCF := &ChannelCmdFactory{
		Signer: signer,
	}

	cmd := signconfigtxCmd(mockCF)

	AddFlags(cmd)

	args := []string{"-f", configtxFile}
	cmd.SetArgs(args)

	assert.NoError(t, cmd.Execute())
}

func TestSignConfigtxMissingConfigTxFlag(t *testing.T) {
	InitMSP()
	resetFlags()

	signer, err := common.GetDefaultSigner()
	if err != nil {
		t.Fatalf("Get default signer error: %v", err)
	}

	mockCF := &ChannelCmdFactory{
		Signer: signer,
	}

	cmd := signconfigtxCmd(mockCF)

	AddFlags(cmd)

	cmd.SetArgs([]string{})

	assert.Error(t, cmd.Execute())
}

func TestSignConfigtxChannelMissingConfigTxFile(t *testing.T) {
	InitMSP()
	resetFlags()

	signer, err := common.GetDefaultSigner()
	if err != nil {
		t.Fatalf("Get default signer error: %v", err)
	}

	mockCF := &ChannelCmdFactory{
		Signer: signer,
	}

	cmd := signconfigtxCmd(mockCF)

	AddFlags(cmd)

	args := []string{"-f", "Non-existant"}
	cmd.SetArgs(args)

	assert.Error(t, cmd.Execute())
}

