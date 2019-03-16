
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:49</date>
//</624455925863813120>

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

//ECDSAP256KeyGenopts包含用于生成曲线P-256的ECDSA密钥的选项。
type ECDSAP256KeyGenOpts struct {
	Temporary bool
}

//算法返回密钥生成算法标识符（要使用）。
func (opts *ECDSAP256KeyGenOpts) Algorithm() string {
	return ECDSAP256
}

//如果要生成的密钥必须是短暂的，则短暂返回true，
//否则为假。
func (opts *ECDSAP256KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

//ECDSAP384密钥包含用于生成曲线P-384的ECDSA密钥的选项。
type ECDSAP384KeyGenOpts struct {
	Temporary bool
}

//算法返回密钥生成算法标识符（要使用）。
func (opts *ECDSAP384KeyGenOpts) Algorithm() string {
	return ECDSAP384
}

//如果要生成的密钥必须是短暂的，则短暂返回true，
//否则为假。
func (opts *ECDSAP384KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

