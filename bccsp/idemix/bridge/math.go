
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:50</date>
//</624455929143758848>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/

package bridge

import (
	"github.com/hyperledger/fabric-amcl/amcl/FP256BN"
	"github.com/hyperledger/fabric/idemix"
)

//大封装AMCL大整数
type Big struct {
	E *FP256BN.BIG
}

func (b *Big) Bytes() ([]byte, error) {
	return idemix.BigToBytes(b.E), nil
}

//ECP封装AMCL椭圆曲线点
type Ecp struct {
	E *FP256BN.ECP
}

func (o *Ecp) Bytes() ([]byte, error) {
	var res []byte
	res = append(res, idemix.BigToBytes(o.E.GetX())...)
	res = append(res, idemix.BigToBytes(o.E.GetY())...)

	return res, nil
}

