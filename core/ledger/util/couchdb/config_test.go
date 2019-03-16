
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456035158986752>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package couchdb

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGetCouchDBDefinition(t *testing.T) {
	expectedAddress := viper.GetString("ledger.state.couchDBConfig.couchDBAddress")

	couchDBDef := GetCouchDBDefinition()
	assert.Equal(t, expectedAddress, couchDBDef.URL)
	assert.Equal(t, "", couchDBDef.Username)
	assert.Equal(t, "", couchDBDef.Password)
	assert.Equal(t, 3, couchDBDef.MaxRetries)
	assert.Equal(t, 20, couchDBDef.MaxRetriesOnStartup)
	assert.Equal(t, time.Second*35, couchDBDef.RequestTimeout)
}

