
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455976744914944>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package resources

//用于ACL检查的结构资源。注意一些支票
//例如lscc_安装是“对等的”（当前对等中的访问检查是
//基于本地MSP）。这些当前不在资源或默认范围内
//提供者
const (
//LSCC资源
	Lscc_Install                   = "lscc/Install"
	Lscc_Deploy                    = "lscc/Deploy"
	Lscc_Upgrade                   = "lscc/Upgrade"
	Lscc_ChaincodeExists           = "lscc/ChaincodeExists"
	Lscc_GetDeploymentSpec         = "lscc/GetDeploymentSpec"
	Lscc_GetChaincodeData          = "lscc/GetChaincodeData"
	Lscc_GetInstantiatedChaincodes = "lscc/GetInstantiatedChaincodes"
	Lscc_GetInstalledChaincodes    = "lscc/GetInstalledChaincodes"
	Lscc_GetCollectionsConfig      = "lscc/GetCollectionsConfig"

//QSCC资源
	Qscc_GetChainInfo       = "qscc/GetChainInfo"
	Qscc_GetBlockByNumber   = "qscc/GetBlockByNumber"
	Qscc_GetBlockByHash     = "qscc/GetBlockByHash"
	Qscc_GetTransactionByID = "qscc/GetTransactionByID"
	Qscc_GetBlockByTxID     = "qscc/GetBlockByTxID"

//CSCC资源
	Cscc_JoinChain                = "cscc/JoinChain"
	Cscc_GetConfigBlock           = "cscc/GetConfigBlock"
	Cscc_GetChannels              = "cscc/GetChannels"
	Cscc_GetConfigTree            = "cscc/GetConfigTree"
	Cscc_SimulateConfigTreeUpdate = "cscc/SimulateConfigTreeUpdate"

//对等资源
	Peer_Propose              = "peer/Propose"
	Peer_ChaincodeToChaincode = "peer/ChaincodeToChaincode"

//事件
	Event_Block         = "event/Block"
	Event_FilteredBlock = "event/FilteredBlock"

//令牌资源
	Token_Issue    = "token/Issue"
	Token_Transfer = "token/Transfer"
	Token_List     = "token/List"
)

