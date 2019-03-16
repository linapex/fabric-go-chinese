
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455992226091008>

package entities

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignedMessage(t *testing.T) {
	ent, err := GetEncrypterSignerEntityForTest("TEST")
	assert.NoError(t, err)
	assert.NotNil(t, ent)

	m := &SignedMessage{Payload: []byte("message"), ID: []byte(ent.ID())}

	err = m.Sign(ent)
	assert.NoError(t, err)
	v, err := m.Verify(ent)
	assert.NoError(t, err)
	assert.True(t, v)
}

func TestSignedMessageErr(t *testing.T) {
	ent, err := GetEncrypterSignerEntityForTest("TEST")
	assert.NoError(t, err)
	assert.NotNil(t, ent)

	m := &SignedMessage{Payload: []byte("message"), ID: []byte(ent.ID())}

	err = m.Sign(nil)
	assert.Error(t, err)
	_, err = m.Verify(nil)
	assert.Error(t, err)

	m = &SignedMessage{Payload: []byte("message"), Sig: []byte("barf")}
	_, err = m.Verify(nil)
	assert.Error(t, err)
}

func TestSignedMessageMarshaller(t *testing.T) {
	m1 := &SignedMessage{Payload: []byte("message"), Sig: []byte("sig"), ID: []byte("ID")}
	m2 := &SignedMessage{}
	b, err := m1.ToBytes()
	assert.NoError(t, err)
	err = m2.FromBytes(b)
	assert.NoError(t, err)
	assert.True(t, reflect.DeepEqual(m1, m2))
}

