
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:30</date>
//</624456098455228416>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package etcdraft

import (
	"testing"

	"github.com/hyperledger/fabric/core/comm"
	"github.com/hyperledger/fabric/orderer/common/cluster"
	"github.com/hyperledger/fabric/orderer/common/localconfig"
	"github.com/hyperledger/fabric/orderer/common/multichannel"
	"github.com/stretchr/testify/assert"
)

func TestNewEtcdRaftConsenter(t *testing.T) {
	srv, err := comm.NewGRPCServer("127.0.0.1:0", comm.ServerConfig{})
	assert.NoError(t, err)
	defer srv.Stop()
	dialer := &cluster.PredicateDialer{}
	consenter := New(dialer, &localconfig.TopLevel{}, comm.ServerConfig{
		SecOpts: &comm.SecureOptions{
			Certificate: []byte{1, 2, 3},
		},
	}, srv, &multichannel.Registrar{})

//声明GRPC服务器的证书已传递给同意者
	assert.Equal(t, []byte{1, 2, 3}, consenter.Cert)
//声明已填充同意者的所有依赖项
	assert.NotNil(t, consenter.Communication)
	assert.NotNil(t, consenter.Chains)
	assert.NotNil(t, consenter.ChainSelector)
	assert.NotNil(t, consenter.Dispatcher)
	assert.NotNil(t, consenter.Logger)
}

