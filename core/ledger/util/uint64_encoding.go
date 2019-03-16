
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456036106899456>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package util

import (
	"encoding/binary"
	"math"

	"github.com/golang/protobuf/proto"
)

//EncodeReverseOrderVaruint64返回uint64数字的字节表示，以便
//首先从maxuint64中减去数字，然后减去所有前导的0xff字节。
//被修剪并替换为此类修剪的字节数。这有助于减小尺寸。
//在字节顺序比较中，此编码确保encodeReverseOrderVaruint64（a）>encodeReverseOrderVaruint64（b），
//如果B> A
func EncodeReverseOrderVarUint64(number uint64) []byte {
	bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(bytes, math.MaxUint64-number)
	numFFBytes := 0
	for _, b := range bytes {
		if b != 0xff {
			break
		}
		numFFBytes++
	}
	size := 8 - numFFBytes
	encodedBytes := make([]byte, size+1)
	encodedBytes[0] = proto.EncodeVarint(uint64(numFFBytes))[0]
	copy(encodedBytes[1:], bytes[numFFBytes:])
	return encodedBytes
}

//decodeReverseOrderVaruint64解码从函数“encodeReverseOrderVaruint64”获得的字节数。
//另外，返回进程中使用的字节数
func DecodeReverseOrderVarUint64(bytes []byte) (uint64, int) {
	s, _ := proto.DecodeVarint(bytes)
	numFFBytes := int(s)
	decodedBytes := make([]byte, 8)
	realBytesNum := 8 - numFFBytes
	copy(decodedBytes[numFFBytes:], bytes[1:realBytesNum+1])
	numBytesConsumed := realBytesNum + 1
	for i := 0; i < numFFBytes; i++ {
		decodedBytes[i] = 0xff
	}
	return (math.MaxUint64 - binary.BigEndian.Uint64(decodedBytes)), numBytesConsumed
}

