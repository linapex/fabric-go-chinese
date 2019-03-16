
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:12</date>
//</624456024635478016>

/*
版权所有IBM公司。保留所有权利。
SPDX许可证标识符：Apache-2.0
**/


package statecouchdb

import (
	"sync"
)

//批处理在单独的goroutine中执行。
type batch interface {
	execute() error
}

//executeBatches在单独的goroutine中执行每个批，如果
//任何批在执行期间返回错误
func executeBatches(batches []batch) error {
	logger.Debugf("Executing batches = %s", batches)
	numBatches := len(batches)
	if numBatches == 0 {
		return nil
	}
	if numBatches == 1 {
		return batches[0].execute()
	}
	var batchWG sync.WaitGroup
	batchWG.Add(numBatches)
	errsChan := make(chan error, numBatches)
	defer close(errsChan)
	for _, b := range batches {
		go func(b batch) {
			defer batchWG.Done()
			if err := b.execute(); err != nil {
				errsChan <- err
				return
			}
		}(b)
	}
	batchWG.Wait()
	if len(errsChan) > 0 {
		return <-errsChan
	}
	return nil
}

