
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:10</date>
//</624456012790763520>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package validation

import "github.com/hyperledger/fabric/core/handlers/validation/api"

//功能定义验证的功能
//在验证事务时应考虑
type Capabilities interface {
	validation.Dependency
//如果此通道中存在所需的未知功能，则SUPPORTED返回错误
	Supported() error

//ForbidDuplicateXdinBlock指定是否允许两个具有相同TxID的事务
//在同一个块中，或者是否将第二个块标记为txvalidationcode_duplicate_txid
	ForbidDuplicateTXIdInBlock() bool

//如果对等端在通道配置中支持ACL，则acls返回true
	ACLs() bool

//如果启用了对专用通道数据（也称为集合）的支持，则private channel data返回true。
	PrivateChannelData() bool

//如果将此通道配置为允许更新到
//现有集合或通过chaincode升级添加新集合（如v1.2所述）
	CollectionUpgrade() bool

//v11validation返回true是否将此通道配置为执行更严格的验证
//事务数（如v1.1中介绍的）。
	V1_1Validation() bool

//v12validation返回true是否将此通道配置为执行更严格的验证
//事务数（如v1.2所述）。
	V1_2Validation() bool

//v13validation如果此通道支持事务验证，则返回true
//如v1.3所述。这包括：
//-如FAB-8812所述，以分类帐键粒度表示的政策
//-新的链码生命周期，如FAB-11237所述
	V1_3Validation() bool

//MetadataLifecycle指示对等端是否应使用已弃用和有问题的
//v1.0/v1.1生命周期，或者是否应该使用更新的每通道对等本地链代码
//计划使用Fabric v1.2发布的元数据包方法
	MetadataLifecycle() bool

//如果此通道支持认可，则keyLevelndOrsement返回true
//如FAB-8812所述，以分类帐键粒度表示的策略
	KeyLevelEndorsement() bool

//如果支持fabric token函数，fabtoken返回true。
	FabToken() bool
}

