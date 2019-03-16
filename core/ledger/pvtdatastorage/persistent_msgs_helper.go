
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456034152353792>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package pvtdatastorage

func newExpiryData() *ExpiryData {
	return &ExpiryData{Map: make(map[string]*Collections)}
}

func (e *ExpiryData) getOrCreateCollections(ns string) *Collections {
	collections, ok := e.Map[ns]
	if !ok {
		collections = &Collections{
			Map:            make(map[string]*TxNums),
			MissingDataMap: make(map[string]bool)}
		e.Map[ns] = collections
	} else {
//由于Protobuf编码/解码，以前
//由于长度为0，初始化的映射现在可能为零。
//因此，我们需要重新初始化映射。
		if collections.Map == nil {
			collections.Map = make(map[string]*TxNums)
		}
		if collections.MissingDataMap == nil {
			collections.MissingDataMap = make(map[string]bool)
		}
	}
	return collections
}

func (e *ExpiryData) addPresentData(ns, coll string, txNum uint64) {
	collections := e.getOrCreateCollections(ns)

	txNums, ok := collections.Map[coll]
	if !ok {
		txNums = &TxNums{}
		collections.Map[coll] = txNums
	}
	txNums.List = append(txNums.List, txNum)
}

func (e *ExpiryData) addMissingData(ns, coll string) {
	collections := e.getOrCreateCollections(ns)
	collections.MissingDataMap[coll] = true
}

func newCollElgInfo(nsCollMap map[string][]string) *CollElgInfo {
	m := &CollElgInfo{NsCollMap: map[string]*CollNames{}}
	for ns, colls := range nsCollMap {
		collNames, ok := m.NsCollMap[ns]
		if !ok {
			collNames = &CollNames{}
			m.NsCollMap[ns] = collNames
		}
		collNames.Entries = colls
	}
	return m
}

