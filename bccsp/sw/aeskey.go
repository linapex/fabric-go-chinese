
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:51</date>
//</624455935481352192>

/*
版权所有IBM Corp.2016保留所有权利。

根据Apache许可证2.0版（以下简称“许可证”）获得许可；
除非符合许可证，否则您不能使用此文件。
您可以在以下网址获得许可证副本：

   http://www.apache.org/licenses/license-2.0

除非适用法律要求或书面同意，软件
根据许可证分发是按“原样”分发的，
无任何明示或暗示的保证或条件。
有关管理权限和
许可证限制。
**/

package sw

import (
	"crypto/sha256"
	"errors"

	"github.com/hyperledger/fabric/bccsp"
)

type aesPrivateKey struct {
	privKey    []byte
	exportable bool
}

//字节将此键转换为其字节表示形式，
//如果允许此操作。
func (k *aesPrivateKey) Bytes() (raw []byte, err error) {
	if k.exportable {
		return k.privKey, nil
	}

	return nil, errors.New("Not supported.")
}

//ski返回此密钥的主题密钥标识符。
func (k *aesPrivateKey) SKI() (ski []byte) {
	hash := sha256.New()
	hash.Write([]byte{0x01})
	hash.Write(k.privKey)
	return hash.Sum(nil)
}

//如果此密钥是对称密钥，则对称返回true，
//如果此密钥是非对称的，则为false
func (k *aesPrivateKey) Symmetric() bool {
	return true
}

//如果此密钥是私钥，则private返回true，
//否则为假。
func (k *aesPrivateKey) Private() bool {
	return true
}

//public key返回非对称公钥/私钥对的相应公钥部分。
//此方法返回对称密钥方案中的错误。
func (k *aesPrivateKey) PublicKey() (bccsp.Key, error) {
	return nil, errors.New("Cannot call this method on a symmetric key.")
}

