
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:59</date>
//</624455968192729088>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/


package semaphore_test

import (
	"context"
	"testing"

	"github.com/hyperledger/fabric/common/semaphore"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestNewSemaphorePanic(t *testing.T) {
	assert.PanicsWithValue(t, "count must be greater than 0", func() { semaphore.New(0) })
}

func TestSemaphoreBlocking(t *testing.T) {
	gt := NewGomegaWithT(t)

	sema := semaphore.New(5)
	for i := 0; i < 5; i++ {
		err := sema.Acquire(context.Background())
		gt.Expect(err).NotTo(HaveOccurred())
	}

	done := make(chan struct{})
	go func() {
		err := sema.Acquire(context.Background())
		gt.Expect(err).NotTo(HaveOccurred())

		close(done)
		sema.Release()
	}()

	gt.Consistently(done).ShouldNot(BeClosed())
	sema.Release()
	gt.Eventually(done).Should(BeClosed())
}

func TestSemaphoreContextError(t *testing.T) {
	gt := NewGomegaWithT(t)

	sema := semaphore.New(1)
	err := sema.Acquire(context.Background())
	gt.Expect(err).NotTo(HaveOccurred())

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	errCh := make(chan error, 1)
	go func() { errCh <- sema.Acquire(ctx) }()

	gt.Eventually(errCh).Should(Receive(Equal(context.Canceled)))
}

func TestSemaphoreReleaseTooMany(t *testing.T) {
	sema := semaphore.New(1)
	assert.PanicsWithValue(t, "semaphore buffer is empty", func() { sema.Release() })
}

