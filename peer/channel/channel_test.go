
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:32</date>
//</624456107728834560>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channel

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitCmdFactory(t *testing.T) {
	t.Run("InitCmdFactory() with PeerDeliverRequired and OrdererRequired", func(t *testing.T) {
		cf, err := InitCmdFactory(EndorserRequired, PeerDeliverRequired, OrdererRequired)
		assert.Nil(t, cf)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ERROR - only a single deliver source is currently supported")
	})
}

