
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:03</date>
//</624455984806367232>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package chaincode

import (
	"github.com/golang/protobuf/proto"
	commonledger "github.com/hyperledger/fabric/common/ledger"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type PendingQueryResult struct {
	batch []*pb.QueryResultBytes
}

func (p *PendingQueryResult) Cut() []*pb.QueryResultBytes {
	batch := p.batch
	p.batch = nil
	return batch
}

func (p *PendingQueryResult) Add(queryResult commonledger.QueryResult) error {
	queryResultBytes, err := proto.Marshal(queryResult.(proto.Message))
	if err != nil {
		chaincodeLogger.Errorf("failed to marshal query result: %s", err)
		return err
	}
	p.batch = append(p.batch, &pb.QueryResultBytes{ResultBytes: queryResultBytes})
	return nil
}

func (p *PendingQueryResult) Size() int {
	return len(p.batch)
}

