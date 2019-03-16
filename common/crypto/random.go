
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455949049925632>

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

**/


package crypto

import (
	"crypto/rand"

	"github.com/pkg/errors"
)

const (
//noncosize是默认的noncosize
	NonceSize = 24
)

//GetRandomBytes返回len随机查找字节
func GetRandomBytes(len int) ([]byte, error) {
	key := make([]byte, len)

//TODO:RAND可以填充更少的字节，然后是LEN
	_, err := rand.Read(key)
	if err != nil {
		return nil, errors.Wrap(err, "error getting random bytes")
	}

	return key, nil
}

//getrandomonce返回长度为noncosize的随机字节数组
func GetRandomNonce() ([]byte, error) {
	return GetRandomBytes(NonceSize)
}

