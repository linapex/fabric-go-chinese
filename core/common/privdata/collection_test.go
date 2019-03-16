
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:07</date>
//</624456001671663616>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package privdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildCollectionKVSKey(t *testing.T) {

	chaincodeCollectionKey := BuildCollectionKVSKey("chaincodeKey")
	assert.Equal(t, "chaincodeKey~collection", chaincodeCollectionKey, "collection keys should end in ~collection")
}

func TestIsCollectionConfigKey(t *testing.T) {

	isCollection := IsCollectionConfigKey("chaincodeKey")
	assert.False(t, isCollection, "key without tilda is not a collection key and should have returned false")

	isCollection = IsCollectionConfigKey("chaincodeKey~collection")
	assert.True(t, isCollection, "key with tilda is a collection key and should have returned true")
}

