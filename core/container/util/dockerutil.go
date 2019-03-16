
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:08</date>
//</624456006490918912>

/*
版权所有2016年伦敦证券交易所版权所有。

SPDX许可证标识符：Apache-2.0
**/


package util

import (
	"runtime"
	"strings"

	docker "github.com/fsouza/go-dockerclient"
	"github.com/hyperledger/fabric/common/metadata"
	"github.com/hyperledger/fabric/core/config"
	"github.com/spf13/viper"
)

//NewDockerClient创建Docker客户端
func NewDockerClient() (client *docker.Client, err error) {
	endpoint := viper.GetString("vm.endpoint")
	tlsenabled := viper.GetBool("vm.docker.tls.enabled")
	if tlsenabled {
		cert := config.GetPath("vm.docker.tls.cert.file")
		key := config.GetPath("vm.docker.tls.key.file")
		ca := config.GetPath("vm.docker.tls.ca.file")
		client, err = docker.NewTLSClient(endpoint, cert, key, ca)
	} else {
		client, err = docker.NewClient(endpoint)
	}
	return
}

func ParseDockerfileTemplate(template string) string {
	r := strings.NewReplacer(
		"$(ARCH)", runtime.GOARCH,
		"$(PROJECT_VERSION)", metadata.Version,
		"$(BASE_VERSION)", metadata.BaseVersion,
		"$(DOCKER_NS)", metadata.DockerNamespace,
		"$(BASE_DOCKER_NS)", metadata.BaseDockerNamespace)

	return r.Replace(template)
}

func GetDockerfileFromConfig(path string) string {
	return ParseDockerfileTemplate(viper.GetString(path))
}

