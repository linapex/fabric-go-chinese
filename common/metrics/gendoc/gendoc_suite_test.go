
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455963016957952>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package gendoc_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGendoc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gendoc Suite")
}

func ParseFile(filename string) (*ast.File, error) {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return f, nil
}

