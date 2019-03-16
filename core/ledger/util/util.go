
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456036178202624>

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


package util

import (
	"reflect"
	"sort"

	"github.com/hyperledger/fabric/common/util"
)

//GetSortedKeys按排序顺序返回映射的键。此函数假定键是字符串
func GetSortedKeys(m interface{}) []string {
	mapVal := reflect.ValueOf(m)
	keyVals := mapVal.MapKeys()
	keys := []string{}
	for _, keyVal := range keyVals {
		keys = append(keys, keyVal.String())
	}
	sort.Strings(keys)
	return keys
}

//GetValuesBysortedKeys按映射键的排序顺序返回列表（listptr）中映射（mapptr）的值。
//此函数假定mapptr是指向映射的指针，listptr是指向列表的指针。键的进一步类型
//假设map为字符串，map和list的值类型相同
func GetValuesBySortedKeys(mapPtr interface{}, listPtr interface{}) {
	mapVal := reflect.ValueOf(mapPtr).Elem()
	keyVals := mapVal.MapKeys()
	if len(keyVals) == 0 {
		return
	}
	keys := make(keys, len(keyVals))
	for i, k := range keyVals {
		keys[i] = newKey(k)
	}
	sort.Sort(keys)
	out := reflect.ValueOf(listPtr).Elem()
	for _, k := range keys {
		val := mapVal.MapIndex(k.Value)
		out.Set(reflect.Append(out, val))
	}
}

type key struct {
	reflect.Value
	str string
}

type keys []*key

func newKey(v reflect.Value) *key {
	return &key{v, v.String()}
}

func (keys keys) Len() int {
	return len(keys)
}

func (keys keys) Swap(i, j int) {
	keys[i], keys[j] = keys[j], keys[i]
}

func (keys keys) Less(i, j int) bool {
	return keys[i].str < keys[j].str
}

//ComputeStringHash计算给定字符串的哈希
func ComputeStringHash(input string) []byte {
	return ComputeHash([]byte(input))
}

//computehash计算给定字节的哈希
func ComputeHash(input []byte) []byte {
	return util.ComputeSHA256(input)
}

