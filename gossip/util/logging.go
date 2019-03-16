
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:24</date>
//</624456072660258816>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package util

import (
	"sync"

	"github.com/hyperledger/fabric/common/flogging"
	"go.uber.org/zap/zapcore"
)

//用于记录器初始化的记录器名称。
const (
	ChannelLogger     = "gossip.channel"
	CommLogger        = "gossip.comm"
	DiscoveryLogger   = "gossip.discovery"
	ElectionLogger    = "gossip.election"
	GossipLogger      = "gossip.gossip"
	CommMockLogger    = "gossip.comm.mock"
	PullLogger        = "gossip.pull"
	ServiceLogger     = "gossip.service"
	StateLogger       = "gossip.state"
	PrivateDataLogger = "gossip.privdata"
)

var loggers = make(map[string]Logger)
var lock = sync.Mutex{}
var testMode bool

//defaulttestspec是八卦测试的默认日志记录级别。
var defaultTestSpec = "WARNING"

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Warning(args ...interface{})
	Warningf(format string, args ...interface{})
	IsEnabledFor(l zapcore.Level) bool
}

//GetLogger返回给定八卦记录器名称和peerID的记录器
func GetLogger(name string, peerID string) Logger {
	if peerID != "" && testMode {
		name = name + "#" + peerID
	}

	lock.Lock()
	defer lock.Unlock()

	if lgr, ok := loggers[name]; ok {
		return lgr
	}

//记录器不存在，请创建一个新的记录器
	lgr := flogging.MustGetLogger(name)
	loggers[name] = lgr
	return lgr
}

//SETUPTESTLOGING设置八卦单元测试的默认日志级别
func SetupTestLogging() {
	testMode = true
	flogging.InitFromSpec(defaultTestSpec)
}

