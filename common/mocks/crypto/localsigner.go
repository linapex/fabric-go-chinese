
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455966062022656>

/*
版权所有IBM Corp.2017保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

                 http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
*/


package crypto

import (
	cb "github.com/hyperledger/fabric/protos/common"
)

//
var FakeLocalSigner = &LocalSigner{
	Identity: []byte("IdentityBytes"),
	Nonce:    []byte("NonceValue"),
}

//localsigner是crypto.localsigner的模拟实现
type LocalSigner struct {
	Identity []byte
	Nonce    []byte
}

//sign返回msg，nil
func (ls *LocalSigner) Sign(msg []byte) ([]byte, error) {
	return msg, nil
}

//NewSignatureHeader返回新的签名头，nil
func (ls *LocalSigner) NewSignatureHeader() (*cb.SignatureHeader, error) {
	return &cb.SignatureHeader{
		Creator: ls.Identity,
		Nonce:   ls.Nonce,
	}, nil
}

