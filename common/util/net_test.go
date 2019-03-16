
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455975792807936>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package util

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/peer"
)

type addr struct {
}

func (*addr) Network() string {
	return ""
}

func (*addr) String() string {
	return "1.2.3.4:5000"
}

func TestExtractAddress(t *testing.T) {
	ctx := context.Background()
	assert.Zero(t, ExtractRemoteAddress(ctx))

	ctx = peer.NewContext(ctx, &peer.Peer{
		Addr: &addr{},
	})
	assert.Equal(t, "1.2.3.4:5000", ExtractRemoteAddress(ctx))
}

