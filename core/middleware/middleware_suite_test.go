
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456036534718464>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package middleware_test

import (
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Middleware Suite")
}

//go：生成伪造者-o fakes/http_handler.go--fake name http handler。亨德勒

type httpHandler interface {
	http.Handler
}

