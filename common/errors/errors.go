
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455951226769408>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package errors

//txvalidationError标记错误与
//交易的确认
type TxValidationError interface {
	error
	IsValid() bool
}

//vsccinfo查找失败错误，指示无法
//从LCCC获取VSCC信息
type VSCCInfoLookupFailureError struct {
	Reason string
}

//错误返回导致失败的原因
func (e VSCCInfoLookupFailureError) Error() string {
	return e.Reason
}

//用于标记事务的vscc背书策略错误
//背书策略检查失败
type VSCCEndorsementPolicyError struct {
	Err error
}

func (e *VSCCEndorsementPolicyError) IsValid() bool {
	return e.Err == nil
}

//错误返回导致失败的原因
func (e VSCCEndorsementPolicyError) Error() string {
	return e.Err.Error()
}

//要指示的vsccexecutionfailureerror错误
//尝试执行vscc时失败
//背书政策检查
type VSCCExecutionFailureError struct {
	Err error
}

//错误返回导致失败的原因
func (e VSCCExecutionFailureError) Error() string {
	return e.Err.Error()
}

func (e *VSCCExecutionFailureError) IsValid() bool {
	return e.Err == nil
}

