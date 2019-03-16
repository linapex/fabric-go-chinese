
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455978464579584>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package chaincode

import (
	"github.com/hyperledger/fabric/core/common/ccprovider"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//ccproviderimpl是ccprovider.chaincodeprovider接口的实现
type CCProviderImpl struct {
	cs *ChaincodeSupport
}

func NewProvider(cs *ChaincodeSupport) *CCProviderImpl {
	return &CCProviderImpl{cs: cs}
}

//执行执行给定上下文和规范（调用或部署）的链代码。
func (c *CCProviderImpl) Execute(txParams *ccprovider.TransactionParams, cccid *ccprovider.CCContext, input *pb.ChaincodeInput) (*pb.Response, *pb.ChaincodeEvent, error) {
	return c.cs.Execute(txParams, cccid, input)
}

//executeLegacyInit执行一个不在lscc表中的链代码
func (c *CCProviderImpl) ExecuteLegacyInit(txParams *ccprovider.TransactionParams, cccid *ccprovider.CCContext, spec *pb.ChaincodeDeploymentSpec) (*pb.Response, *pb.ChaincodeEvent, error) {
	return c.cs.ExecuteLegacyInit(txParams, cccid, spec)
}

//
func (c *CCProviderImpl) Stop(ccci *ccprovider.ChaincodeContainerInfo) error {
	return c.cs.Stop(ccci)
}

