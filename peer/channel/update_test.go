
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:33</date>
//</624456109138120704>

/*
版权所有IBM Corp.2017保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

                 http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
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

const mockChannel = "mockChannel"

func TestUpdateChannel(t *testing.T) {
	InitMSP()
	resetFlags()

	dir, err := ioutil.TempDir("/tmp", "createinvaltest-")
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
		BroadcastFactory: mockBroadcastClientFactory,
		Signer:           signer,
		DeliverClient:    &mockDeliverClient{},
	}

	cmd := updateCmd(mockCF)

	AddFlags(cmd)

	args := []string{"-c", mockChannel, "-f", configtxFile, "-o", "localhost:7050"}
	cmd.SetArgs(args)

	assert.NoError(t, cmd.Execute())
}

func TestUpdateChannelMissingConfigTxFlag(t *testing.T) {
	InitMSP()
	resetFlags()

	signer, err := common.GetDefaultSigner()
	if err != nil {
		t.Fatalf("Get default signer error: %v", err)
	}

	mockCF := &ChannelCmdFactory{
		BroadcastFactory: mockBroadcastClientFactory,
		Signer:           signer,
		DeliverClient:    &mockDeliverClient{},
	}

	cmd := updateCmd(mockCF)

	AddFlags(cmd)

	args := []string{"-c", mockChannel, "-o", "localhost:7050"}
	cmd.SetArgs(args)

	assert.Error(t, cmd.Execute())
}

func TestUpdateChannelMissingConfigTxFile(t *testing.T) {
	InitMSP()
	resetFlags()

	signer, err := common.GetDefaultSigner()
	if err != nil {
		t.Fatalf("Get default signer error: %v", err)
	}

	mockCF := &ChannelCmdFactory{
		BroadcastFactory: mockBroadcastClientFactory,
		Signer:           signer,
		DeliverClient:    &mockDeliverClient{},
	}

	cmd := updateCmd(mockCF)

	AddFlags(cmd)

	args := []string{"-c", mockChannel, "-f", "Non-existant", "-o", "localhost:7050"}
	cmd.SetArgs(args)

	assert.Error(t, cmd.Execute())
}

func TestUpdateChannelMissingChannelID(t *testing.T) {
	InitMSP()
	resetFlags()

	dir, err := ioutil.TempDir("/tmp", "createinvaltest-")
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
		BroadcastFactory: mockBroadcastClientFactory,
		Signer:           signer,
		DeliverClient:    &mockDeliverClient{},
	}

	cmd := updateCmd(mockCF)

	AddFlags(cmd)

	args := []string{"-f", configtxFile, "-o", "localhost:7050"}
	cmd.SetArgs(args)

	assert.Error(t, cmd.Execute())
}

