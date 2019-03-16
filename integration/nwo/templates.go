
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456079153041408>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package nwo

//模板可用于提供自定义模板以生成配置树。
type Templates struct {
	ConfigTx string `yaml:"configtx,omitempty"`
	Core     string `yaml:"core,omitempty"`
	Crypto   string `yaml:"crypto,omitempty"`
	Orderer  string `yaml:"orderer,omitempty"`
}

func (t *Templates) ConfigTxTemplate() string {
	if t.ConfigTx != "" {
		return t.ConfigTx
	}
	return DefaultConfigTxTemplate
}

func (t *Templates) CoreTemplate() string {
	if t.Core != "" {
		return t.Core
	}
	return DefaultCoreTemplate
}

func (t *Templates) CryptoTemplate() string {
	if t.Crypto != "" {
		return t.Crypto
	}
	return DefaultCryptoTemplate
}

func (t *Templates) OrdererTemplate() string {
	if t.Orderer != "" {
		return t.Orderer
	}
	return DefaultOrdererTemplate
}

