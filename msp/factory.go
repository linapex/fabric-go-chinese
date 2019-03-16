
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:26</date>
//</624456081719955456>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package msp

import (
	"github.com/pkg/errors"
)

type MSPVersion int

const (
	MSPv1_0 = iota
	MSPv1_1
	MSPv1_3
)

//Newopts表示
type NewOpts interface {
//GetVersion返回要实例化的MSP版本
	GetVersion() MSPVersion
}

//newbaseopts是所有msp实例化opts的默认基类型
type NewBaseOpts struct {
	Version MSPVersion
}

func (o *NewBaseOpts) GetVersion() MSPVersion {
	return o.Version
}

//bccspnewopts包含实例化新的基于bccsp的（x509）msp的选项
type BCCSPNewOpts struct {
	NewBaseOpts
}

//idemixnewopts包含实例化新的基于idemix的MSP的选项
type IdemixNewOpts struct {
	NewBaseOpts
}

//新建根据传递的opt创建新的msp实例
func New(opts NewOpts) (MSP, error) {
	switch opts.(type) {
	case *BCCSPNewOpts:
		switch opts.GetVersion() {
		case MSPv1_0:
			return newBccspMsp(MSPv1_0)
		case MSPv1_1:
			return newBccspMsp(MSPv1_1)
		case MSPv1_3:
			return newBccspMsp(MSPv1_3)
		default:
			return nil, errors.Errorf("Invalid *BCCSPNewOpts. Version not recognized [%v]", opts.GetVersion())
		}
	case *IdemixNewOpts:
		switch opts.GetVersion() {
		case MSPv1_3:
			return newIdemixMsp(MSPv1_3)
		case MSPv1_1:
			return newIdemixMsp(MSPv1_1)
		default:
			return nil, errors.Errorf("Invalid *IdemixNewOpts. Version not recognized [%v]", opts.GetVersion())
		}
	default:
		return nil, errors.Errorf("Invalid msp.NewOpts instance. It must be either *BCCSPNewOpts or *IdemixNewOpts. It was [%v]", opts)
	}
}

