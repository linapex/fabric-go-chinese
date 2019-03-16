
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:53</date>
//</624455942129324032>

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

package attrmgr_test

import (
	"crypto/x509"
	"testing"

	"github.com/hyperledger/fabric/common/attrmgr"
	"github.com/stretchr/testify/assert"
)

//测试ttrs测试属性
func TestAttrs(t *testing.T) {
	mgr := attrmgr.New()
	attrs := []attrmgr.Attribute{
		&Attribute{Name: "attr1", Value: "val1"},
		&Attribute{Name: "attr2", Value: "val2"},
		&Attribute{Name: "attr3", Value: "val3"},
		&Attribute{Name: "boolAttr", Value: "true"},
	}
	reqs := []attrmgr.AttributeRequest{
		&AttributeRequest{Name: "attr1", Require: false},
		&AttributeRequest{Name: "attr2", Require: true},
		&AttributeRequest{Name: "boolAttr", Require: true},
		&AttributeRequest{Name: "noattr1", Require: false},
	}
	cert := &x509.Certificate{}

//验证证书是否没有属性
	at, err := mgr.GetAttributesFromCert(cert)
	if err != nil {
		t.Fatalf("Failed to GetAttributesFromCert: %s", err)
	}
	numAttrs := len(at.Names())
	assert.True(t, numAttrs == 0, "expecting 0 attributes but found %d", numAttrs)

//向证书添加属性
	err = mgr.ProcessAttributeRequestsForCert(reqs, attrs, cert)
	if err != nil {
		t.Fatalf("Failed to ProcessAttributeRequestsForCert: %s", err)
	}

//从证书中获取属性并验证计数是否正确
	at, err = mgr.GetAttributesFromCert(cert)
	if err != nil {
		t.Fatalf("Failed to GetAttributesFromCert: %s", err)
	}
	numAttrs = len(at.Names())
	assert.True(t, numAttrs == 3, "expecting 3 attributes but found %d", numAttrs)

//检查单个属性
	checkAttr(t, "attr1", "val1", at)
	checkAttr(t, "attr2", "val2", at)
	checkAttr(t, "attr3", "", at)
	checkAttr(t, "noattr1", "", at)
	assert.NoError(t, at.True("boolAttr"))

//否定测试用例：添加不存在的必需属性
	reqs = []attrmgr.AttributeRequest{
		&AttributeRequest{Name: "noattr1", Require: true},
	}
	err = mgr.ProcessAttributeRequestsForCert(reqs, attrs, cert)
	assert.Error(t, err)
}

func checkAttr(t *testing.T, name, val string, attrs *attrmgr.Attributes) {
	v, ok, err := attrs.Value(name)
	assert.NoError(t, err)
	if val == "" {
		assert.False(t, attrs.Contains(name), "contains attribute '%s'", name)
		assert.False(t, ok, "attribute '%s' was found", name)
	} else {
		assert.True(t, attrs.Contains(name), "does not contain attribute '%s'", name)
		assert.True(t, ok, "attribute '%s' was not found", name)
		assert.True(t, v == val, "incorrect value for '%s'; expected '%s' but found '%s'", name, val, v)
	}
}

type Attribute struct {
	Name, Value string
}

func (a *Attribute) GetName() string {
	return a.Name
}

func (a *Attribute) GetValue() string {
	return a.Value
}

type AttributeRequest struct {
	Name    string
	Require bool
}

func (ar *AttributeRequest) GetName() string {
	return ar.Name
}

func (ar *AttributeRequest) IsRequired() bool {
	return ar.Require
}

