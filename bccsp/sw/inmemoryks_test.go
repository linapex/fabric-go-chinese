
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:51</date>
//</624455937029050368>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package sw

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidStore(t *testing.T) {
	t.Parallel()

	ks := NewInMemoryKeyStore()

	err := ks.StoreKey(nil)
	assert.EqualError(t, err, "key is nil")
}

func TestInvalidLoad(t *testing.T) {
	t.Parallel()

	ks := NewInMemoryKeyStore()

	_, err := ks.GetKey(nil)
	assert.EqualError(t, err, "ski is nil or empty")
}

func TestNoKeyFound(t *testing.T) {
	t.Parallel()

	ks := NewInMemoryKeyStore()

	ski := []byte("foo")
	_, err := ks.GetKey(ski)
	assert.EqualError(t, err, fmt.Sprintf("no key found for ski %x", ski))
}

func TestStoreLoad(t *testing.T) {
	t.Parallel()

	ks := NewInMemoryKeyStore()

//为要查找的密钥库生成密钥
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	cspKey := &ecdsaPrivateKey{privKey}

//存储密钥
	err = ks.StoreKey(cspKey)
	assert.NoError(t, err)

//加载键
	key, err := ks.GetKey(cspKey.SKI())
	assert.NoError(t, err)

	assert.Equal(t, cspKey, key)
}

func TestReadOnly(t *testing.T) {
	t.Parallel()
	ks := NewInMemoryKeyStore()
	readonly := ks.ReadOnly()
	assert.Equal(t, false, readonly)
}

func TestStoreExisting(t *testing.T) {
	t.Parallel()

	ks := NewInMemoryKeyStore()

//为要查找的密钥库生成密钥
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	assert.NoError(t, err)
	cspKey := &ecdsaPrivateKey{privKey}

//存储密钥
	err = ks.StoreKey(cspKey)
	assert.NoError(t, err)

//再次存储密钥
	err = ks.StoreKey(cspKey)
	assert.EqualError(t, err, fmt.Sprintf("ski %x already exists in the keystore", cspKey.SKI()))
}

