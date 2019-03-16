
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456077588566016>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package commands

type Generate struct {
	Config string
	Output string
}

func (c Generate) SessionName() string {
	return "cryptogen-generate"
}

func (c Generate) Args() []string {
	return []string{
		"generate",
		"--config", c.Config,
		"--output", c.Output,
	}
}

