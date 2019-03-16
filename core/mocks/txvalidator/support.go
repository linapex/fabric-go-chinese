
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456037369384960>

/*
版权所有IBM Corp.2017保留所有权利。

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


package support

import (
	"sync"

	"github.com/hyperledger/fabric/common/channelconfig"
	mockpolicies "github.com/hyperledger/fabric/common/mocks/policies"
	"github.com/hyperledger/fabric/common/policies"
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/msp"
	"github.com/hyperledger/fabric/protos/common"
)

type Support struct {
	LedgerVal     ledger.PeerLedger
	MSPManagerVal msp.MSPManager
	ApplyVal      error
	ACVal         channelconfig.ApplicationCapabilities

	sync.Mutex
	capabilitiesInvokeCount int
	mspManagerInvokeCount   int
}

func (ms *Support) Capabilities() channelconfig.ApplicationCapabilities {
	ms.Lock()
	defer ms.Unlock()
	ms.capabilitiesInvokeCount++
	return ms.ACVal
}

//分类帐返回分类帐
func (ms *Support) Ledger() ledger.PeerLedger {
	return ms.LedgerVal
}

//mspmanager返回mspmanagerval
func (ms *Support) MSPManager() msp.MSPManager {
	ms.Lock()
	defer ms.Unlock()
	ms.mspManagerInvokeCount++
	return ms.MSPManagerVal
}

//应用返回ApplyVal
func (ms *Support) Apply(configtx *common.ConfigEnvelope) error {
	return ms.ApplyVal
}

func (ms *Support) PolicyManager() policies.Manager {
	return &mockpolicies.Manager{}
}

func (ms *Support) GetMSPIDs(cid string) []string {
	return []string{"SampleOrg"}
}

func (ms *Support) CapabilitiesInvokeCount() int {
	ms.Lock()
	defer ms.Unlock()
	return ms.capabilitiesInvokeCount
}

func (ms *Support) MSPManagerInvokeCount() int {
	ms.Lock()
	defer ms.Unlock()
	return ms.mspManagerInvokeCount
}

