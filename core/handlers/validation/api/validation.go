
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:10</date>
//</624456013071781888>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package validation

import "github.com/hyperledger/fabric/protos/common"

//参数定义用于验证的参数
type Argument interface {
	Dependency
//arg返回参数的字节数
	Arg() []byte
}

//依赖项标记传递给init（）方法的依赖项
type Dependency interface{}

//ContextDatum定义从验证器传递的附加数据
//进入validate（）调用
type ContextDatum interface{}

//插件验证事务
type Plugin interface {
//如果在事务内部给定位置的操作，则validate返回nil
//在给定的位置，在给定的块中是有效的，否则是错误的。
	Validate(block *common.Block, namespace string, txPosition int, actionPosition int, contextData ...ContextDatum) error

//init将依赖项插入插件的实例中
	Init(dependencies ...Dependency) error
}

//PluginFactory创建插件的新实例
type PluginFactory interface {
	New() Plugin
}

//ExecutionFailureError指示验证
//由于执行问题而失败，因此
//无法计算事务验证状态
type ExecutionFailureError struct {
	Reason string
}

//错误表示这是一个错误，并且还包含
//错误的原因
func (e *ExecutionFailureError) Error() string {
	return e.Reason
}

