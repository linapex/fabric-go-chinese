
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:09</date>
//</624456011163373568>

/*
版权所有SecureKey Technologies Inc.保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package decoration

import (
	"encoding/binary"
	"testing"

	"github.com/hyperledger/fabric/protos/peer"
	"github.com/stretchr/testify/assert"
)

const (
	decorationKey = "sequence"
)

func TestApplyDecorations(t *testing.T) {
	iterations := 15
	decorators := createNDecorators(iterations)
	initialInput := &peer.ChaincodeInput{Decorations: make(map[string][]byte)}
	seq := make([]byte, 4)
	binary.BigEndian.PutUint32(seq, 0)
	initialInput.Decorations[decorationKey] = seq

	finalInput := Apply(nil, initialInput, decorators...)
	for i := 0; i < iterations; i++ {
		assert.Equal(t, uint32(i), decorators[i].(*mockDecorator).sequence,
			"Expected decorators to be applied in the provided sequence")
	}

	assert.Equal(t, uint32(iterations), binary.BigEndian.Uint32(finalInput.Decorations[decorationKey]),
		"Expected decorators to be applied in the provided sequence")
}

func createNDecorators(n int) []Decorator {
	decorators := make([]Decorator, n)
	for i := 0; i < n; i++ {
		decorators[i] = &mockDecorator{}
	}
	return decorators
}

type mockDecorator struct {
	sequence uint32
}

func (d *mockDecorator) Decorate(proposal *peer.Proposal,
	input *peer.ChaincodeInput) *peer.ChaincodeInput {

	d.sequence = binary.BigEndian.Uint32(input.Decorations[decorationKey])
	binary.BigEndian.PutUint32(input.Decorations[decorationKey], d.sequence+1)

	return input
}

