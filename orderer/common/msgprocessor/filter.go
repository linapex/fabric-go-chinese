
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:29</date>
//</624456094088957952>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package msgprocessor

import (
	"errors"

	ab "github.com/hyperledger/fabric/protos/common"
)

//拒绝时，空消息筛选器返回errEmptyMessage。
var ErrEmptyMessage = errors.New("Message was empty")

//规则定义一个过滤函数，它接受、拒绝或转发（到下一个规则）信封。
type Rule interface {
//应用将规则应用于给定信封，成功或返回错误
	Apply(message *ab.Envelope) error
}

//EmptyRejectRule拒绝空消息
var EmptyRejectRule = Rule(emptyRejectRule{})

type emptyRejectRule struct{}

func (a emptyRejectRule) Apply(message *ab.Envelope) error {
	if message.Payload == nil {
		return ErrEmptyMessage
	}
	return nil
}

//AcceptRule始终返回Accept作为Apply的结果
var AcceptRule = Rule(acceptRule{})

type acceptRule struct{}

func (a acceptRule) Apply(message *ab.Envelope) error {
	return nil
}

//规则集用于应用规则集合
type RuleSet struct {
	rules []Rule
}

//new ruleset使用给定的规则有序列表创建新规则集
func NewRuleSet(rules []Rule) *RuleSet {
	return &RuleSet{
		rules: rules,
	}
}

//应用按顺序应用为此集指定的规则，在有效时返回nil，在无效时返回err。
func (rs *RuleSet) Apply(message *ab.Envelope) error {
	for _, rule := range rs.rules {
		err := rule.Apply(message)
		if err != nil {
			return err
		}
	}
	return nil
}

