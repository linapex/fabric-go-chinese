
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:13</date>
//</624456026774573056>

/*
版权所有IBM公司。保留所有权利。
SPDX许可证标识符：Apache-2.0
**/


package lockbasedtxmgr

import (
	"github.com/hyperledger/fabric/core/ledger"
	"github.com/hyperledger/fabric/protos/common"
)

//collnamevalidator验证命名空间中是否存在集合
//这将在模拟器/查询执行器的上下文中实例化。
type collNameValidator struct {
	ccInfoProvider ledger.DeployedChaincodeInfoProvider
	queryExecutor  *lockBasedQueryExecutor
	cache          collConfigCache
}

func newCollNameValidator(ccInfoProvider ledger.DeployedChaincodeInfoProvider, qe *lockBasedQueryExecutor) *collNameValidator {
	return &collNameValidator{ccInfoProvider, qe, make(collConfigCache)}
}

func (v *collNameValidator) validateCollName(ns, coll string) error {
	if !v.cache.isPopulatedFor(ns) {
		conf, err := v.retrieveCollConfigFromStateDB(ns)
		if err != nil {
			return err
		}
		v.cache.populate(ns, conf)
	}
	if !v.cache.containsCollName(ns, coll) {
		return &ledger.InvalidCollNameError{
			Ns:   ns,
			Coll: coll,
		}
	}
	return nil
}

func (v *collNameValidator) retrieveCollConfigFromStateDB(ns string) (*common.CollectionConfigPackage, error) {
	logger.Debugf("retrieveCollConfigFromStateDB() begin - ns=[%s]", ns)
	ccInfo, err := v.ccInfoProvider.ChaincodeInfo(ns, v.queryExecutor)
	if err != nil {
		return nil, err
	}
	if ccInfo == nil || ccInfo.CollectionConfigPkg == nil {
		return nil, &ledger.CollConfigNotDefinedError{Ns: ns}
	}
	confPkg := ccInfo.CollectionConfigPkg
	logger.Debugf("retrieveCollConfigFromStateDB() successfully retrieved - ns=[%s], confPkg=[%s]", ns, confPkg)
	return confPkg, nil
}

type collConfigCache map[collConfigkey]bool

type collConfigkey struct {
	ns, coll string
}

func (c collConfigCache) populate(ns string, pkg *common.CollectionConfigPackage) {
//具有空集合名称的条目，用于指示为命名空间“ns”填充了缓存。
//请参见函数“ispopulatedfor”
	c[collConfigkey{ns, ""}] = true
	for _, config := range pkg.Config {
		sConfig := config.GetStaticCollectionConfig()
		if sConfig == nil {
			continue
		}
		c[collConfigkey{ns, sConfig.Name}] = true
	}
}

func (c collConfigCache) isPopulatedFor(ns string) bool {
	return c[collConfigkey{ns, ""}]
}

func (c collConfigCache) containsCollName(ns, coll string) bool {
	return c[collConfigkey{ns, coll}]
}

