
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456036270477312>

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
**/


package util

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSortedKeys(t *testing.T) {
	mapKeyValue := make(map[string]int)
	mapKeyValue["blue"] = 10
	mapKeyValue["apple"] = 15
	mapKeyValue["red"] = 12
	mapKeyValue["123"] = 22
	mapKeyValue["a"] = 33
	mapKeyValue[""] = 30
	assert.Equal(t, []string{"", "123", "a", "apple", "blue", "red"}, GetSortedKeys(mapKeyValue))
}

func TestGetValuesBySortedKeys(t *testing.T) {
	type name struct {
		fName string
		lName string
	}
	mapKeyValue := make(map[string]*name)
	mapKeyValue["2"] = &name{"Two", "two"}
	mapKeyValue["3"] = &name{"Three", "three"}
	mapKeyValue["5"] = &name{"Five", "five"}
	mapKeyValue[""] = &name{"None", "none"}

	sortedRes := []*name{}
	GetValuesBySortedKeys(&mapKeyValue, &sortedRes)
	assert.Equal(
		t,
		[]*name{{"None", "none"}, {"Two", "two"}, {"Three", "three"}, {"Five", "five"}},
		sortedRes,
	)
}

func TestBasicEncodingDecoding(t *testing.T) {
	for i := 0; i < 10000; i++ {
		value := EncodeReverseOrderVarUint64(uint64(i))
		nextValue := EncodeReverseOrderVarUint64(uint64(i + 1))
		if !(bytes.Compare(value, nextValue) > 0) {
			t.Fatalf("A smaller integer should result into greater bytes. Encoded bytes for [%d] is [%x] and for [%d] is [%x]",
				i, i+1, value, nextValue)
		}
		decodedValue, _ := DecodeReverseOrderVarUint64(value)
		if decodedValue != uint64(i) {
			t.Fatalf("Value not same after decoding. Original value = [%d], decode value = [%d]", i, decodedValue)
		}
	}
}

func TestDecodingAppendedValues(t *testing.T) {
	appendedValues := []byte{}
	for i := 0; i < 1000; i++ {
		appendedValues = append(appendedValues, EncodeReverseOrderVarUint64(uint64(i))...)
	}

	len := 0
	value := uint64(0)
	for i := 0; i < 1000; i++ {
		appendedValues = appendedValues[len:]
		value, len = DecodeReverseOrderVarUint64(appendedValues)
		if value != uint64(i) {
			t.Fatalf("expected value = [%d], decode value = [%d]", i, value)
		}
	}
}

