
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455951159660544>

/*
版权所有IBM公司保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package diag_test

import (
	"testing"

	"github.com/hyperledger/fabric/common/diag"
	"github.com/hyperledger/fabric/common/flogging/floggingtest"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

func TestCaptureGoRoutines(t *testing.T) {
	gt := NewGomegaWithT(t)
	output, err := diag.CaptureGoRoutines()
	gt.Expect(err).NotTo(HaveOccurred())

	gt.Expect(output).To(MatchRegexp(`goroutine \d+ \[running\]:`))
	gt.Expect(output).To(ContainSubstring("github.com/hyperledger/fabric/common/diag.CaptureGoRoutines"))
}

func TestLogGoRoutines(t *testing.T) {
	gt := NewGomegaWithT(t)
	logger, recorder := floggingtest.NewTestLogger(t, floggingtest.Named("goroutine"))
	diag.LogGoRoutines(logger)

	gt.Expect(recorder).To(gbytes.Say(`goroutine \d+ \[running\]:`))
}

