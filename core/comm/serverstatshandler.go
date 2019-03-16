
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455995501842432>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package comm

import (
	"context"

	"github.com/hyperledger/fabric/common/metrics"
	"google.golang.org/grpc/stats"
)

type ServerStatsHandler struct {
	OpenConnCounter   metrics.Counter
	ClosedConnCounter metrics.Counter
}

func (h *ServerStatsHandler) TagRPC(ctx context.Context, info *stats.RPCTagInfo) context.Context {
	return ctx
}

func (h *ServerStatsHandler) HandleRPC(ctx context.Context, s stats.RPCStats) {}

func (h *ServerStatsHandler) TagConn(ctx context.Context, info *stats.ConnTagInfo) context.Context {
	return ctx
}

func (h *ServerStatsHandler) HandleConn(ctx context.Context, s stats.ConnStats) {
	switch s.(type) {
	case *stats.ConnBegin:
		h.OpenConnCounter.Add(1)
	case *stats.ConnEnd:
		h.ClosedConnCounter.Add(1)
	}
}

