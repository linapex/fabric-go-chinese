
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:11</date>
//</624456019862360064>

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


package kvledger

import "github.com/hyperledger/fabric/core/ledger"

type recoverable interface {
//是否需要恢复，是否需要恢复。
//如果需要恢复，此方法还返回从中开始恢复的块号。
//LastavailableBlock是已提交到块存储的最大块号
	ShouldRecover(lastAvailableBlock uint64) (bool, uint64, error)
//commitListBlock重新提交块
	CommitLostBlock(block *ledger.BlockAndPvtData) error
}

type recoverer struct {
	firstBlockNum uint64
	recoverable   recoverable
}

