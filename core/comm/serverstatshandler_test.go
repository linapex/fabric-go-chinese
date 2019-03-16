
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455995568951296>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package comm_test

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/hyperledger/fabric/common/metrics"
	"github.com/hyperledger/fabric/common/metrics/metricsfakes"
	"github.com/hyperledger/fabric/core/comm"
	testpb "github.com/hyperledger/fabric/core/comm/testdata/grpc"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc"
	"google.golang.org/grpc/stats"
)

func TestConnectionCounters(t *testing.T) {
	t.Parallel()
	gt := NewGomegaWithT(t)

	openConn := &metricsfakes.Counter{}
	closedConn := &metricsfakes.Counter{}
	sh := &comm.ServerStatsHandler{
		OpenConnCounter:   openConn,
		ClosedConnCounter: closedConn,
	}

	for i := 1; i <= 10; i++ {
		sh.HandleConn(context.Background(), &stats.ConnBegin{})
		gt.Expect(openConn.AddCallCount()).To(Equal(i))
	}

	for i := 1; i <= 5; i++ {
		sh.HandleConn(context.Background(), &stats.ConnEnd{})
		gt.Expect(closedConn.AddCallCount()).To(Equal(i))
	}
}

func TestConnMetricsGRPCServer(t *testing.T) {
	t.Parallel()
	gt := NewGomegaWithT(t)

	openConn := &metricsfakes.Counter{}
	closedConn := &metricsfakes.Counter{}
	fakeProvider := &metricsfakes.Provider{}
	fakeProvider.NewCounterStub = func(o metrics.CounterOpts) metrics.Counter {
		switch o.Name {
		case "conn_opened":
			return openConn
		case "conn_closed":
			return closedConn
		default:
			panic("unknown counter")
		}
	}

	listener, err := net.Listen("tcp", "localhost:0")
	gt.Expect(err).NotTo(HaveOccurred())
	srv, err := comm.NewGRPCServerFromListener(
		listener,
		comm.ServerConfig{
			SecOpts:         &comm.SecureOptions{UseTLS: false},
			MetricsProvider: fakeProvider,
		},
	)
	gt.Expect(err).NotTo(HaveOccurred())

//注册GRPC测试服务器
	testpb.RegisterEmptyServiceServer(srv.Server(), &emptyServiceServer{})

//启动服务器
	go srv.Start()
	defer srv.Stop()

//测试GRPC连接计数
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	gt.Expect(openConn.AddCallCount()).To(Equal(0))
	gt.Expect(closedConn.AddCallCount()).To(Equal(0))

//创建GRPC客户端连接
	var clientConns []*grpc.ClientConn
	for i := 1; i <= 3; i++ {
		clientConn, err := grpc.DialContext(ctx, listener.Addr().String(), grpc.WithInsecure())
		gt.Expect(err).NotTo(HaveOccurred())
		clientConns = append(clientConns, clientConn)

//调用服务
		client := testpb.NewEmptyServiceClient(clientConn)
		_, err = client.EmptyCall(context.Background(), &testpb.Empty{})
		gt.Expect(err).NotTo(HaveOccurred())
		gt.Expect(openConn.AddCallCount()).To(Equal(i))
	}

	for i, conn := range clientConns {
		gt.Expect(closedConn.AddCallCount()).Should(Equal(i))
		conn.Close()
		gt.Eventually(closedConn.AddCallCount, time.Second).Should(Equal(i + 1))
	}
}

