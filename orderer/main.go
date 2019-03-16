
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:31</date>
//</624456102989271040>

/*
版权所有IBM Corp.2017保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


//package main是订购方二进制文件的入口点
//并且只调用server.main（）函数。没有其他
//此包中应包含函数。
package main

import "github.com/hyperledger/fabric/orderer/common/server"

func main() {
	server.Main()
}

