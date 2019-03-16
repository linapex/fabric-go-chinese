
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455962719162368>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/hyperledger/fabric/common/metrics/gendoc"
	"golang.org/x/tools/go/packages"
)

//
//
//文档。

var templatePath = flag.String(
	"template",
	"docs/source/metrics_reference.rst.tmpl",
	"The documentation template.",
)

func main() {
	flag.Parse()

	patterns := flag.Args()
	if len(patterns) == 0 {
		patterns = []string{"github.com/hyperledger/fabric/..."}
	}

	pkgs, err := packages.Load(&packages.Config{Mode: packages.LoadSyntax}, patterns...)
	if err != nil {
		panic(err)
	}

	options, err := gendoc.Options(pkgs)
	if err != nil {
		panic(err)
	}

	cells, err := gendoc.NewCells(options)
	if err != nil {
		panic(err)
	}

	funcMap := template.FuncMap{
		"PrometheusTable": func() string {
			buf := &bytes.Buffer{}
			gendoc.NewPrometheusTable(cells).Generate(buf)
			return buf.String()
		},
		"StatsdTable": func() string {
			buf := &bytes.Buffer{}
			gendoc.NewStatsdTable(cells).Generate(buf)
			return buf.String()
		},
	}

	docTemplate, err := ioutil.ReadFile(*templatePath)
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("metrics_reference").Funcs(funcMap).Parse(string(docTemplate))
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(os.Stdout, ""); err != nil {
		panic(err)
	}
}

