
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:29</date>
//</624456092964884480>

/*
版权所有IBM Corp.2017保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package cluster

import (
	"context"
	"io"

	"github.com/hyperledger/fabric/common/flogging"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/protos/orderer"
	"google.golang.org/grpc"
)

//去：生成mokery-dir。-name dispatcher-case underline-output./mocks/

//调度员发送请求
type Dispatcher interface {
	DispatchSubmit(ctx context.Context, request *orderer.SubmitRequest) (*orderer.SubmitResponse, error)
	DispatchStep(ctx context.Context, request *orderer.StepRequest) (*orderer.StepResponse, error)
}

//去：生成mokery-dir。-name submitstream-case underline-output./mocks/

//SubmitStream定义用于发送的GRPC流
//交易，并接收相应的响应
type SubmitStream interface {
	Send(response *orderer.SubmitResponse) error
	Recv() (*orderer.SubmitRequest, error)
	grpc.ServerStream
}

//服务定义筏式服务
type Service struct {
	Dispatcher Dispatcher
	Logger     *flogging.FabricLogger
	StepLogger *flogging.FabricLogger
}

//将消息转发到此服务器中的raft fsm
func (s *Service) Step(ctx context.Context, request *orderer.StepRequest) (*orderer.StepResponse, error) {
	addr := util.ExtractRemoteAddress(ctx)
	s.StepLogger.Debugf("Connection from %s", addr)
	defer s.StepLogger.Debugf("Closing connection from %s", addr)
	response, err := s.Dispatcher.DispatchStep(ctx, request)
	if err != nil {
		s.Logger.Warningf("Handling of Step() from %s failed: %+v", addr, err)
	}
	return response, err
}

//提交接受交易记录
func (s *Service) Submit(stream orderer.Cluster_SubmitServer) error {
	addr := util.ExtractRemoteAddress(stream.Context())
	s.Logger.Debugf("Connection from %s", addr)
	defer s.Logger.Debugf("Closing connection from %s", addr)
	for {
		err := s.handleSubmit(stream, addr)
		if err == io.EOF {
			s.Logger.Debugf("%s disconnected", addr)
			return nil
		}
		if err != nil {
			return err
		}
//否则，不会发生错误，因此我们继续下一个迭代
	}
}

func (s *Service) handleSubmit(stream SubmitStream, addr string) error {
	request, err := stream.Recv()
	if err == io.EOF {
		return err
	}
	if err != nil {
		s.Logger.Warningf("Stream read from %s failed: %v", addr, err)
		return err
	}
	response, err := s.Dispatcher.DispatchSubmit(stream.Context(), request)
	if err != nil {
		s.Logger.Warningf("Handling of Propose() from %s failed: %+v", addr, err)
		return err
	}
	err = stream.Send(response)
	if err != nil {
		s.Logger.Warningf("Send() failed: %v", err)
	}
	return err
}

