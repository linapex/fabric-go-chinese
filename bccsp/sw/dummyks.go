
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:51</date>
//</624455935728816128>

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
	"errors"

	"github.com/hyperledger/fabric/bccsp"
)

//NewDummyKeyStore实例化虚拟密钥存储
//既不加载也不存储密钥
func NewDummyKeyStore() bccsp.KeyStore {
	return &dummyKeyStore{}
}

//dummykeystore是一个既不加载也不存储密钥的只读密钥库。
type dummyKeyStore struct {
}

//如果此密钥库是只读的，则read only返回true，否则返回false。
//如果readonly为true，则storekey将失败。
func (ks *dummyKeyStore) ReadOnly() bool {
	return true
}

//getkey返回一个key对象，其ski是通过的。
func (ks *dummyKeyStore) GetKey(ski []byte) (bccsp.Key, error) {
	return nil, errors.New("Key not found. This is a dummy KeyStore")
}

//storekey将密钥k存储在此密钥库中。
//如果此密钥库是只读的，则该方法将失败。
func (ks *dummyKeyStore) StoreKey(k bccsp.Key) error {
	return errors.New("Cannot store key. This is a dummy read-only KeyStore")
}

