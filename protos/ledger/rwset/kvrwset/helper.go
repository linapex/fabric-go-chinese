
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:34</date>
//</624456117350567936>

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


package kvrwset

import (
	"bytes"
)

//setrawreads将“readsinfo”字段设置为查询执行的原始kvreads
func (rqi *RangeQueryInfo) SetRawReads(kvReads []*KVRead) {
	rqi.ReadsInfo = &RangeQueryInfo_RawReads{
		RawReads: &QueryReads{
			KvReads: kvReads,
		},
	}
}

//setmerkelsummary将“readsinfo”字段设置为查询结果的原始kvards的merkle摘要
func (rqi *RangeQueryInfo) SetMerkelSummary(merkleSummary *QueryReadsMerkleSummary) {
	rqi.ReadsInfo = &RangeQueryInfo_ReadsMerkleHashes{merkleSummary}
}

//equal验证give merklesummary是否等于此
func (ms *QueryReadsMerkleSummary) Equal(anotherMS *QueryReadsMerkleSummary) bool {
	if anotherMS == nil {
		return false
	}
	if ms.MaxDegree != anotherMS.MaxDegree ||
		ms.MaxLevel != anotherMS.MaxLevel ||
		len(ms.MaxLevelHashes) != len(anotherMS.MaxLevelHashes) {
		return false
	}
	for i := 0; i < len(ms.MaxLevelHashes); i++ {
		if !bytes.Equal(ms.MaxLevelHashes[i], anotherMS.MaxLevelHashes[i]) {
			return false
		}
	}
	return true
}

