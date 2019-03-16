
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:22</date>
//</624456065752240128>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package comm

import (
	"sync"

	proto "github.com/hyperledger/fabric/protos/gossip"
	"github.com/pkg/errors"
)

//ReceivedMessageImpl是ReceivedMessage的实现
type ReceivedMessageImpl struct {
	*proto.SignedGossipMessage
	lock     sync.Locker
	conn     *connection
	connInfo *proto.ConnectionInfo
}

//GetSourceEnvelope返回接收到的消息所在的信封
//建筑用
func (m *ReceivedMessageImpl) GetSourceEnvelope() *proto.Envelope {
	return m.Envelope
}

//Respond sends a msg to the source that sent the ReceivedMessageImpl
func (m *ReceivedMessageImpl) Respond(msg *proto.GossipMessage) {
	sMsg, err := msg.NoopSign()
	if err != nil {
		err = errors.WithStack(err)
		m.conn.logger.Errorf("Failed creating SignedGossipMessage: %+v", err)
		return
	}
	m.conn.send(sMsg, func(e error) {}, blockingSend)
}

//GetGossipMessage返回内部GossipMessage
func (m *ReceivedMessageImpl) GetGossipMessage() *proto.SignedGossipMessage {
	return m.SignedGossipMessage
}

//getConnectionInfo返回有关远程对等机的信息
//发出信息
func (m *ReceivedMessageImpl) GetConnectionInfo() *proto.ConnectionInfo {
	return m.connInfo
}

//ACK向发送方返回消息确认
func (m *ReceivedMessageImpl) Ack(err error) {
	ackMsg := &proto.GossipMessage{
		Nonce: m.GetGossipMessage().Nonce,
		Content: &proto.GossipMessage_Ack{
			Ack: &proto.Acknowledgement{},
		},
	}
	if err != nil {
		ackMsg.GetAck().Error = err.Error()
	}
	m.Respond(ackMsg)
}

