
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:24</date>
//</624456073905967104>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package idemix

import (
	"github.com/hyperledger/fabric-amcl/amcl/FP256BN"
	"github.com/pkg/errors"
)

//nonrevokedprove是处理撤销的zk证明系统的验证器。
type nonRevocationVerifier interface {
//recomputefscontribution重新计算非撤销证明对zkp挑战的贡献
	recomputeFSContribution(proof *NonRevocationProof, chal *FP256BN.BIG, epochPK *FP256BN.ECP2, proofSRh *FP256BN.BIG) ([]byte, error)
}

//noNonReversionVerifier是一个空的非ReversionVerifier，它产生一个空的贡献
type nopNonRevocationVerifier struct{}

func (verifier *nopNonRevocationVerifier) recomputeFSContribution(proof *NonRevocationProof, chal *FP256BN.BIG, epochPK *FP256BN.ECP2, proofSRh *FP256BN.BIG) ([]byte, error) {
	return nil, nil
}

//GetUnrevocationVerifier返回绑定到传递的吊销算法的UnrevocationVerifier
func getNonRevocationVerifier(algorithm RevocationAlgorithm) (nonRevocationVerifier, error) {
	switch algorithm {
	case ALG_NO_REVOCATION:
		return &nopNonRevocationVerifier{}, nil
	default:
//未知的吊销算法
		return nil, errors.Errorf("unknown revocation algorithm %d", algorithm)
	}
}

