
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:05</date>
//</624455992377085952>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package statebased

import "fmt"

//背书政策标识的角色类型
type RoleType string

const (
//roletypember标识组织的成员标识
	RoleTypeMember = RoleType("MEMBER")
//roletypeer标识组织的对等身份
	RoleTypePeer = RoleType("PEER")
)

//
//
//
type RoleTypeDoesNotExistError struct {
	RoleType RoleType
}

func (r *RoleTypeDoesNotExistError) Error() string {
	return fmt.Sprintf("role type %s does not exist", r.RoleType)
}

//
//
//
//
type KeyEndorsementPolicy interface {
//
	Policy() ([]byte, error)

//
//
//在第一个参数中指定。在其他方面，期望的角色
//取决于通道的配置：如果它支持节点OU，则为
//
//如果不是的话。
	AddOrgs(roleType RoleType, organizations ...string) error

//
//
	DelOrgs(organizations ...string)

//
	ListOrgs() []string
}

