
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455994650398720>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package comm

import "github.com/hyperledger/fabric/common/metrics"

var (
	openConnCounterOpts = metrics.CounterOpts{
		Namespace: "grpc",
		Subsystem: "comm",
		Name:      "conn_opened",
		Help:      "gRPC connections opened. Open minus closed is the active number of connections.",
	}

	closedConnCounterOpts = metrics.CounterOpts{
		Namespace: "grpc",
		Subsystem: "comm",
		Name:      "conn_closed",
		Help:      "gRPC connections closed. Open minus closed is the active number of connections.",
	}
)

func NewServerStatsHandler(p metrics.Provider) *ServerStatsHandler {
	return &ServerStatsHandler{
		OpenConnCounter:   p.NewCounter(openConnCounterOpts),
		ClosedConnCounter: p.NewCounter(closedConnCounterOpts),
	}
}

