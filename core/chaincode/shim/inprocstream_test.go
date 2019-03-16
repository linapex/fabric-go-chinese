
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455992779739136>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package shim

import (
	"testing"

	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
)

func TestRecvChannelClosedError(t *testing.T) {
	ch := make(chan *pb.ChaincodeMessage)

	stream := newInProcStream(ch, ch)

//关闭频道
	close(ch)

//尝试调用关闭的接收通道应返回错误
	_, err := stream.Recv()
	if assert.Error(t, err, "Should return an error") {
		assert.Contains(t, err.Error(), "channel is closed")
	}
}

