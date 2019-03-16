
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:30</date>
//</624456097087885312>

//+建设！窗户

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package server

import (
	"os"
	"syscall"

	"github.com/hyperledger/fabric/common/diag"
)

func addPlatformSignals(sigs map[os.Signal]func()) map[os.Signal]func() {
	sigs[syscall.SIGUSR1] = func() { diag.LogGoRoutines(logger.Named("diag")) }
	return sigs
}

