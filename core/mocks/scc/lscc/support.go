
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456037256138752>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package lscc

import (
	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/protos/common"
	"github.com/hyperledger/fabric/protos/peer"
)

type MockSupport struct {
	PutChaincodeToLocalStorageErr    error
	GetChaincodeFromLocalStorageRv   ccprovider.CCPackage
	GetChaincodeFromLocalStorageErr  error
	GetChaincodesFromLocalStorageRv  *peer.ChaincodeQueryResponse
	GetChaincodesFromLocalStorageErr error
	GetInstantiationPolicyRv         []byte
	GetInstantiationPolicyErr        error
	CheckInstantiationPolicyErr      error
	GetInstantiationPolicyMap        map[string][]byte
	CheckInstantiationPolicyMap      map[string]error
	CheckCollectionConfigErr         error
}

func (s *MockSupport) PutChaincodeToLocalStorage(ccpack ccprovider.CCPackage) error {
	return s.PutChaincodeToLocalStorageErr
}

func (s *MockSupport) GetChaincodeFromLocalStorage(ccname string, ccversion string) (ccprovider.CCPackage, error) {
	return s.GetChaincodeFromLocalStorageRv, s.GetChaincodeFromLocalStorageErr
}

func (s *MockSupport) GetChaincodesFromLocalStorage() (*peer.ChaincodeQueryResponse, error) {
	return s.GetChaincodesFromLocalStorageRv, s.GetChaincodesFromLocalStorageErr
}

func (s *MockSupport) GetInstantiationPolicy(channel string, ccpack ccprovider.CCPackage) ([]byte, error) {
	if s.GetInstantiationPolicyMap != nil {
		str := ccpack.GetChaincodeData().Name + ccpack.GetChaincodeData().Version
		s.GetInstantiationPolicyMap[str] = []byte(str)
		return []byte(ccpack.GetChaincodeData().Name + ccpack.GetChaincodeData().Version), nil
	}
	return s.GetInstantiationPolicyRv, s.GetInstantiationPolicyErr
}

func (s *MockSupport) CheckInstantiationPolicy(signedProp *peer.SignedProposal, chainName string, instantiationPolicy []byte) error {
	if s.CheckInstantiationPolicyMap != nil {
		return s.CheckInstantiationPolicyMap[string(instantiationPolicy)]
	}
	return s.CheckInstantiationPolicyErr
}

func (s *MockSupport) CheckCollectionConfig(collectionConfig *common.CollectionConfig, channelName string) error {
	return s.CheckCollectionConfigErr
}

