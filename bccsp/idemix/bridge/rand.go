
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:50</date>
//</624455929311531008>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package bridge

import (
	"github.com/hyperledger/fabric-amcl/amcl"
	cryptolib "github.com/hyperledger/fabric/idemix"
)

//newrandorpanic返回新的amcl prg或panic
func NewRandOrPanic() *amcl.RAND {
	rng, err := cryptolib.GetRand()
	if err != nil {
		panic(err)
	}
	return rng
}

