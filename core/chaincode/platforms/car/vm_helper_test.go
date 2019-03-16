
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:04</date>
//</624455987838849024>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package car_test

import (
	"bytes"
	"fmt"

	"github.com/fsouza/go-dockerclient"
	"github.com/hyperledger/fabric/core/chaincode/platforms"
	"github.com/hyperledger/fabric/core/chaincode/platforms/car"
	"github.com/hyperledger/fabric/core/container"
	cutil "github.com/hyperledger/fabric/core/container/util"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//虚拟机实现虚拟机管理功能。
type VM struct {
	Client *docker.Client
}

//new vm创建一个新的vm实例。
func NewVM() (*VM, error) {
	client, err := cutil.NewDockerClient()
	if err != nil {
		return nil, err
	}
	VM := &VM{Client: client}
	return VM, nil
}

//BuildChainCodeContainer为提供的链代码规范生成容器
func (vm *VM) BuildChaincodeContainer(spec *pb.ChaincodeSpec) error {
	codePackage, err := container.GetChaincodePackageBytes(platforms.NewRegistry(&car.Platform{}), spec)
	if err != nil {
		return fmt.Errorf("Error getting chaincode package bytes: %s", err)
	}

	cds := &pb.ChaincodeDeploymentSpec{ChaincodeSpec: spec, CodePackage: codePackage}
	dockerSpec, err := platforms.NewRegistry(&car.Platform{}).GenerateDockerBuild(
		cds.CCType(),
		cds.Path(),
		cds.Name(),
		cds.Version(),
		cds.Bytes(),
	)
	if err != nil {
		return fmt.Errorf("Error getting chaincode docker image: %s", err)
	}

	output := bytes.NewBuffer(nil)

	err = vm.Client.BuildImage(docker.BuildImageOptions{
		Name:         spec.ChaincodeId.Name,
		InputStream:  dockerSpec,
		OutputStream: output,
	})
	if err != nil {
		return fmt.Errorf("Error building docker: %s (output = %s)", err, output.String())
	}

	return nil
}

