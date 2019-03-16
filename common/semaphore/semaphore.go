
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455968092065792>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package semaphore

import "context"

type Semaphore chan struct{}

func New(count int) Semaphore {
	if count <= 0 {
		panic("count must be greater than 0")
	}
	return make(chan struct{}, count)
}

func (s Semaphore) Acquire(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case s <- struct{}{}:
		return nil
	}
}

func (s Semaphore) Release() {
	select {
	case <-s:
	default:
		panic("semaphore buffer is empty")
	}
}

