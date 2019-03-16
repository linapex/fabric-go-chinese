
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456077374656512>

/*
版权所有IBM公司保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package helpers

import (
	"encoding/base32"
	"fmt"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/hyperledger/fabric/common/util"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func AssertImagesExist(imageNames ...string) {
	dockerClient, err := docker.NewClientFromEnv()
	Expect(err).NotTo(HaveOccurred())

	for _, imageName := range imageNames {
		images, err := dockerClient.ListImages(docker.ListImagesOptions{
			Filter: imageName,
		})
		ExpectWithOffset(1, err).NotTo(HaveOccurred())

		if len(images) != 1 {
			Fail(fmt.Sprintf("missing required image: %s", imageName), 1)
		}
	}
}

//uniquename为容器名称生成base-32 enocded uuid。
func UniqueName() string {
	name := base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(util.GenerateBytesUUID())
	return strings.ToLower(name)
}

