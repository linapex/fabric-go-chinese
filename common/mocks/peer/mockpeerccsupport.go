
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455966548561920>

/*
版权所有IBM Corp.2017保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/


package peer

import (
	"fmt"

	pb "github.com/hyperledger/fabric/protos/peer"
)

//
type MockPeerCCSupport struct {
	ccStream map[string]*MockCCComm
}

//
func NewMockPeerSupport() *MockPeerCCSupport {
	return &MockPeerCCSupport{ccStream: make(map[string]*MockCCComm)}
}

//addcc将cc添加到mockpeerccsupport
func (mp *MockPeerCCSupport) AddCC(name string, recv chan *pb.ChaincodeMessage, send chan *pb.ChaincodeMessage) (*MockCCComm, error) {
	if mp.ccStream[name] != nil {
		return nil, fmt.Errorf("CC %s already added", name)
	}
	mcc := &MockCCComm{name: name, recvStream: recv, sendStream: send}
	mp.ccStream[name] = mcc
	return mcc, nil
}

//
func (mp *MockPeerCCSupport) GetCC(name string) (*MockCCComm, error) {
	s := mp.ccStream[name]
	if s == nil {
		return nil, fmt.Errorf("CC %s not added", name)
	}
	return s, nil
}

//
func (mp *MockPeerCCSupport) GetCCMirror(name string) *MockCCComm {
	s := mp.ccStream[name]
	if s == nil {
		return nil
	}

	return &MockCCComm{name: name, recvStream: s.sendStream, sendStream: s.recvStream, skipClose: true}
}

//
func (mp *MockPeerCCSupport) RemoveCC(name string) error {
	if mp.ccStream[name] == nil {
		return fmt.Errorf("CC %s not added", name)
	}
	delete(mp.ccStream, name)
	return nil
}

//
func (mp *MockPeerCCSupport) RemoveAll() error {
	mp.ccStream = make(map[string]*MockCCComm)
	return nil
}

