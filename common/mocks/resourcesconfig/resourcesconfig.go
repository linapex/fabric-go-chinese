
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455966749888512>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package resourceconfig

type MockChaincodeDefinition struct {
	NameRv          string
	VersionRv       string
	EndorsementStr  string
	ValidationStr   string
	ValidationBytes []byte
	HashRv          []byte
}

func (m *MockChaincodeDefinition) CCName() string {
	return m.NameRv
}

func (m *MockChaincodeDefinition) Hash() []byte {
	return m.HashRv
}

func (m *MockChaincodeDefinition) CCVersion() string {
	return m.VersionRv
}

func (m *MockChaincodeDefinition) Validation() (string, []byte) {
	return m.ValidationStr, m.ValidationBytes
}

func (m *MockChaincodeDefinition) Endorsement() string {
	return m.EndorsementStr
}

