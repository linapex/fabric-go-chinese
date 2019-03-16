
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:17</date>
//</624456042939420672>

//+构建插件已启用，CGO
//+构建达尔文，go1.10 linux，go1.10 linux，go1.9，！PPC64

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package scc

//createpluginssccs创建所有由plugin加载的系统链码。
func CreatePluginSysCCs(p *Provider) []SelfDescribingSysCC {
	var sdscs []SelfDescribingSysCC
	for _, pscc := range loadSysCCs(p) {
		sdscs = append(sdscs, &SysCCWrapper{SCC: pscc})
	}
	return sdscs
}

