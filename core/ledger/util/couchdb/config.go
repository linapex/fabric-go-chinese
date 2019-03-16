
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456035083489280>

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


package couchdb

import (
	"time"

	"github.com/spf13/viper"
)

//couchdbdef包含参数
type CouchDBDef struct {
	URL                   string
	Username              string
	Password              string
	MaxRetries            int
	MaxRetriesOnStartup   int
	RequestTimeout        time.Duration
	CreateGlobalChangesDB bool
}

//GetCouchDBDefinition exposes the useCouchDB variable
func GetCouchDBDefinition() *CouchDBDef {

	couchDBAddress := viper.GetString("ledger.state.couchDBConfig.couchDBAddress")
	username := viper.GetString("ledger.state.couchDBConfig.username")
	password := viper.GetString("ledger.state.couchDBConfig.password")
	maxRetries := viper.GetInt("ledger.state.couchDBConfig.maxRetries")
	maxRetriesOnStartup := viper.GetInt("ledger.state.couchDBConfig.maxRetriesOnStartup")
	requestTimeout := viper.GetDuration("ledger.state.couchDBConfig.requestTimeout")
	createGlobalChangesDB := viper.GetBool("ledger.state.couchDBConfig.createGlobalChangesDB")

	return &CouchDBDef{couchDBAddress, username, password, maxRetries, maxRetriesOnStartup, requestTimeout, createGlobalChangesDB}
}

