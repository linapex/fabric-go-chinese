
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:26</date>
//</624456080189034496>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package runner

import (
	"time"

	"github.com/hyperledger/fabric/integration/helpers"
)

const DefaultStartTimeout = 30 * time.Second

//defaultnamer是默认的命名函数。
var DefaultNamer NameFunc = helpers.UniqueName

//namefunc用于生成容器名称。
type NameFunc func() string

