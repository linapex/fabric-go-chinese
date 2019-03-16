
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456036949954560>

/*
版权所有IBM Corp.2018保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package ccprovider

import (
	commonledger "github.com/hyperledger/fabric/common/ledger"
	"github.com/hyperledger/fabric/core/common/ccprovider"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/peer"
)

type ExecuteChaincodeResultProvider interface {
	ExecuteChaincodeResult() (*peer.Response, *peer.ChaincodeEvent, error)
}

//MockcProviderFactory是返回的工厂
//ccprovider.chaincodeprovider接口的模拟实现
type MockCcProviderFactory struct {
	ExecuteResultProvider ExecuteChaincodeResultProvider
}

//NewChaincodeProvider返回ccProvider.ChaincodeProvider接口的模拟实现
func (c *MockCcProviderFactory) NewChaincodeProvider() ccprovider.ChaincodeProvider {
	return &MockCcProviderImpl{ExecuteResultProvider: c.ExecuteResultProvider}
}

//mockccproviderimpl是chaincode提供程序的模拟实现
type MockCcProviderImpl struct {
	ExecuteResultProvider    ExecuteChaincodeResultProvider
	ExecuteChaincodeResponse *peer.Response
}

type MockTxSim struct {
	GetTxSimulationResultsRv *ledger.TxSimulationResults
}

func (m *MockTxSim) GetState(namespace string, key string) ([]byte, error) {
	return nil, nil
}

func (m *MockTxSim) GetStateMultipleKeys(namespace string, keys []string) ([][]byte, error) {
	return nil, nil
}

func (m *MockTxSim) GetStateRangeScanIterator(namespace string, startKey string, endKey string) (commonledger.ResultsIterator, error) {
	return nil, nil
}

func (m *MockTxSim) GetStateRangeScanIteratorWithMetadata(namespace string, startKey, endKey string, metadata map[string]interface{}) (ledger.QueryResultsIterator, error) {
	return nil, nil
}

func (m *MockTxSim) ExecuteQuery(namespace, query string) (commonledger.ResultsIterator, error) {
	return nil, nil
}

func (m *MockTxSim) ExecuteQueryWithMetadata(namespace, query string, metadata map[string]interface{}) (ledger.QueryResultsIterator, error) {
	return nil, nil
}

func (m *MockTxSim) Done() {
}

func (m *MockTxSim) SetState(namespace string, key string, value []byte) error {
	return nil
}

func (m *MockTxSim) DeleteState(namespace string, key string) error {
	return nil
}

func (m *MockTxSim) SetStateMultipleKeys(namespace string, kvs map[string][]byte) error {
	return nil
}

func (m *MockTxSim) ExecuteUpdate(query string) error {
	return nil
}

func (m *MockTxSim) GetTxSimulationResults() (*ledger.TxSimulationResults, error) {
	return m.GetTxSimulationResultsRv, nil
}

func (m *MockTxSim) DeletePrivateData(namespace, collection, key string) error {
	return nil
}

func (m *MockTxSim) ExecuteQueryOnPrivateData(namespace, collection, query string) (commonledger.ResultsIterator, error) {
	return nil, nil
}

func (m *MockTxSim) GetPrivateData(namespace, collection, key string) ([]byte, error) {
	return nil, nil
}

func (m *MockTxSim) GetPrivateDataMultipleKeys(namespace, collection string, keys []string) ([][]byte, error) {
	return nil, nil
}

func (m *MockTxSim) GetPrivateDataRangeScanIterator(namespace, collection, startKey, endKey string) (commonledger.ResultsIterator, error) {
	return nil, nil
}

func (m *MockTxSim) SetPrivateData(namespace, collection, key string, value []byte) error {
	return nil
}

func (m *MockTxSim) SetPrivateDataMultipleKeys(namespace, collection string, kvs map[string][]byte) error {
	return nil
}

func (m *MockTxSim) GetStateMetadata(namespace, key string) (map[string][]byte, error) {
	return nil, nil
}

func (m *MockTxSim) GetPrivateDataMetadata(namespace, collection, key string) (map[string][]byte, error) {
	return nil, nil
}

func (m *MockTxSim) GetPrivateDataMetadataByHash(namespace, collection string, keyhash []byte) (map[string][]byte, error) {
	return nil, nil
}

func (m *MockTxSim) SetStateMetadata(namespace, key string, metadata map[string][]byte) error {
	return nil
}

func (m *MockTxSim) DeleteStateMetadata(namespace, key string) error {
	return nil
}

func (m *MockTxSim) SetPrivateDataMetadata(namespace, collection, key string, metadata map[string][]byte) error {
	return nil
}

func (m *MockTxSim) DeletePrivateDataMetadata(namespace, collection, key string) error {
	return nil
}

//ExecuteInit执行给定上下文和规范部署的链代码
func (c *MockCcProviderImpl) ExecuteLegacyInit(txParams *ccprovider.TransactionParams, cccid *ccprovider.CCContext, spec *peer.ChaincodeDeploymentSpec) (*peer.Response, *peer.ChaincodeEvent, error) {
	return &peer.Response{}, nil, nil
}

//执行执行给定上下文和规范调用的链代码
func (c *MockCcProviderImpl) Execute(txParams *ccprovider.TransactionParams, cccid *ccprovider.CCContext, spec *peer.ChaincodeInput) (*peer.Response, *peer.ChaincodeEvent, error) {
	return &peer.Response{}, nil, nil
}

//stop停止给定上下文和部署规范的链代码
func (c *MockCcProviderImpl) Stop(ccci *ccprovider.ChaincodeContainerInfo) error {
	return nil
}

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

