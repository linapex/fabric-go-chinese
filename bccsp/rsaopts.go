
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:51</date>
//</624455935137419264>

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


package bccsp

//RSA1024KeyGenopts包含用于在1024安全性下生成RSA密钥的选项。
type RSA1024KeyGenOpts struct {
	Temporary bool
}

//算法返回密钥生成算法标识符（要使用）。
func (opts *RSA1024KeyGenOpts) Algorithm() string {
	return RSA1024
}

//如果要生成的密钥必须是短暂的，则短暂返回true，
//否则为假。
func (opts *RSA1024KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

//RSA2048KeyGenopts包含2048安全级别的RSA密钥生成选项。
type RSA2048KeyGenOpts struct {
	Temporary bool
}

//算法返回密钥生成算法标识符（要使用）。
func (opts *RSA2048KeyGenOpts) Algorithm() string {
	return RSA2048
}

//如果要生成的密钥必须是短暂的，则短暂返回true，
//否则为假。
func (opts *RSA2048KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

//RSA3072KeyGenopts包含用于3072安全性下生成RSA密钥的选项。
type RSA3072KeyGenOpts struct {
	Temporary bool
}

//算法返回密钥生成算法标识符（要使用）。
func (opts *RSA3072KeyGenOpts) Algorithm() string {
	return RSA3072
}

//如果要生成的密钥必须是短暂的，则短暂返回true，
//否则为假。
func (opts *RSA3072KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

//rsa4096keygenopts包含用于在4096安全性下生成RSA密钥的选项。
type RSA4096KeyGenOpts struct {
	Temporary bool
}

//算法返回密钥生成算法标识符（要使用）。
func (opts *RSA4096KeyGenOpts) Algorithm() string {
	return RSA4096
}

//如果要生成的密钥必须是短暂的，则短暂返回true，
//否则为假。
func (opts *RSA4096KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

