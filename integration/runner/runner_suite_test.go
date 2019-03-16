
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:26</date>
//</624456080423915520>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package runner_test

import (
	"testing"

	docker "github.com/fsouza/go-dockerclient"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRunner(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Runner Suite")
}

func ContainerExists(client *docker.Client, name string) func() bool {
	return func() bool {
		_, err := client.InspectContainer(name)
		if err != nil {
			_, ok := err.(*docker.NoSuchContainer)
			return !ok
		}
		return false
	}
}

