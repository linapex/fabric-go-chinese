
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:35</date>
//</624456120257220608>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package peer

//name实现platforms.namedescriber接口
func (cs *ChaincodeSpec) Name() string {
	if cs.ChaincodeId == nil {
		return ""
	}

	return cs.ChaincodeId.Name
}

func (cs *ChaincodeSpec) Version() string {
	if cs.ChaincodeId == nil {
		return ""
	}

	return cs.ChaincodeId.Version
}

//path实现platforms.pathDescriber接口
func (cs *ChaincodeSpec) Path() string {
	if cs.ChaincodeId == nil {
		return ""
	}

	return cs.ChaincodeId.Path
}

func (cs *ChaincodeSpec) CCType() string {
	return cs.Type.String()
}

//path实现platforms.pathDescriber接口
func (cds *ChaincodeDeploymentSpec) Path() string {
	if cds.ChaincodeSpec == nil {
		return ""
	}

	return cds.ChaincodeSpec.Path()
}

//bytes实现platforms.codepackage接口
func (cds *ChaincodeDeploymentSpec) Bytes() []byte {
	return cds.CodePackage
}

func (cds *ChaincodeDeploymentSpec) CCType() string {
	if cds.ChaincodeSpec == nil {
		return ""
	}

	return cds.ChaincodeSpec.CCType()
}

func (cds *ChaincodeDeploymentSpec) Name() string {
	if cds.ChaincodeSpec == nil {
		return ""
	}

	return cds.ChaincodeSpec.Name()
}

func (cds *ChaincodeDeploymentSpec) Version() string {
	if cds.ChaincodeSpec == nil {
		return ""
	}

	return cds.ChaincodeSpec.Version()
}

