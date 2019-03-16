
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:55</date>
//</624455951096745984>

/*
版权所有IBM公司保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package diag

import (
	"bytes"
	"runtime/pprof"
)

type Logger interface {
	Infof(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}

func CaptureGoRoutines() (string, error) {
	var buf bytes.Buffer
	err := pprof.Lookup("goroutine").WriteTo(&buf, 2)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func LogGoRoutines(logger Logger) {
	output, err := CaptureGoRoutines()
	if err != nil {
		logger.Errorf("failed to capture go routines: %s", err)
		return
	}

	logger.Infof("Go routines report:\n%s", output)
}

