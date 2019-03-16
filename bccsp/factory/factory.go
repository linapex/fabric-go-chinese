
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:49</date>
//</624455925947699200>

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

package factory

import (
	"sync"

	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/common/flogging"
	"github.com/pkg/errors"
)

var (
//缺省BCCSP
	defaultBCCSP bccsp.BCCSP

//当尚未调用initfactures时（只应发生
//在测试用例中），临时使用此bccsp
	bootBCCSP bccsp.BCCSP

//BCCSP工厂
	bccspMap map[string]bccsp.BCCSP

//工厂初始化时的同步
	factoriesInitOnce sync.Once
	bootBCCSPInitOnce sync.Once

//工厂初始化错误
	factoriesInitError error

	logger = flogging.MustGetLogger("bccsp")
)

//bccspfactory用于获取bccsp接口的实例。
//一个工厂有一个用来称呼它的名字。
type BCCSPFactory interface {

//name返回此工厂的名称
	Name() string

//get返回使用opts的bccsp实例。
	Get(opts *FactoryOpts) (bccsp.BCCSP, error)
}

//getdefault返回非短暂（长期）bccsp
func GetDefault() bccsp.BCCSP {
	if defaultBCCSP == nil {
		logger.Warning("Before using BCCSP, please call InitFactories(). Falling back to bootBCCSP.")
		bootBCCSPInitOnce.Do(func() {
			var err error
			f := &SWFactory{}
			bootBCCSP, err = f.Get(GetDefaultOpts())
			if err != nil {
				panic("BCCSP Internal error, failed initialization with GetDefaultOpts!")
			}
		})
		return bootBCCSP
	}
	return defaultBCCSP
}

//getbccsp返回根据输入中传递的选项创建的bccsp。
func GetBCCSP(name string) (bccsp.BCCSP, error) {
	csp, ok := bccspMap[name]
	if !ok {
		return nil, errors.Errorf("Could not find BCCSP, no '%s' provider", name)
	}
	return csp, nil
}

func initBCCSP(f BCCSPFactory, config *FactoryOpts) error {
	csp, err := f.Get(config)
	if err != nil {
		return errors.Errorf("Could not initialize BCCSP %s [%s]", f.Name(), err)
	}

	logger.Debugf("Initialize BCCSP [%s]", f.Name())
	bccspMap[f.Name()] = csp
	return nil
}

