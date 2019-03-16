
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:11</date>
//</624456018700537856>

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


package historydb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var strKeySep = string(CompositeKeySep)

func TestConstructCompositeKey(t *testing.T) {
	compositeKey := ConstructCompositeHistoryKey("ns1", "key1", 1, 1)
	assert.NotNil(t, compositeKey)
//HistoryLevelDB_test.go测试实际输出
}

func TestConstructPartialCompositeKey(t *testing.T) {
	compositeStartKey := ConstructPartialCompositeHistoryKey("ns1", "key1", false)
	compositeEndKey := ConstructPartialCompositeHistoryKey("ns1", "key1", true)

	assert.Equal(t, []byte("ns1"+strKeySep+"key1"+strKeySep), compositeStartKey)
	assert.Equal(t, []byte("ns1"+strKeySep+"key1"+strKeySep+string([]byte{0xff})), compositeEndKey)
}

func TestSplitCompositeKey(t *testing.T) {
	compositeFullKey := []byte("ns1" + strKeySep + "key1" + strKeySep + "extra bytes to split")
	compositePartialKey := ConstructPartialCompositeHistoryKey("ns1", "key1", false)

	_, extraBytes := SplitCompositeHistoryKey(compositeFullKey, compositePartialKey)
//第二个位置应保留被拆分的额外字节。
	assert.Equal(t, []byte("extra bytes to split"), extraBytes)
}

