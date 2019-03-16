
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:58</date>
//</624455962421366784>

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
*/


package util

import (
	"bytes"
	"testing"
)

func TestBasicEncodingDecoding(t *testing.T) {
	for i := 0; i < 10000; i++ {
		value := EncodeOrderPreservingVarUint64(uint64(i))
		nextValue := EncodeOrderPreservingVarUint64(uint64(i + 1))
		if !(bytes.Compare(value, nextValue) < 0) {
			t.Fatalf("A smaller integer should result into smaller bytes. Encoded bytes for [%d] is [%x] and for [%d] is [%x]",
				i, i+1, value, nextValue)
		}
		decodedValue, _ := DecodeOrderPreservingVarUint64(value)
		if decodedValue != uint64(i) {
			t.Fatalf("Value not same after decoding. Original value = [%d], decode value = [%d]", i, decodedValue)
		}
	}
}

func TestDecodingAppendedValues(t *testing.T) {
	appendedValues := []byte{}
	for i := 0; i < 1000; i++ {
		appendedValues = append(appendedValues, EncodeOrderPreservingVarUint64(uint64(i))...)
	}

	len := 0
	value := uint64(0)
	for i := 0; i < 1000; i++ {
		appendedValues = appendedValues[len:]
		value, len = DecodeOrderPreservingVarUint64(appendedValues)
		if value != uint64(i) {
			t.Fatalf("expected value = [%d], decode value = [%d]", i, value)
		}
	}
}

