
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:54</date>
//</624455946516566016>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package channelconfig

import (
	"github.com/hyperledger/fabric/common/policies"
)

func LogSanityChecks(res Resources) {
	pm := res.PolicyManager()
	for _, policyName := range []string{policies.ChannelReaders, policies.ChannelWriters} {
		_, ok := pm.GetPolicy(policyName)
		if !ok {
			logger.Warningf("Current configuration has no policy '%s', this will likely cause problems in production systems", policyName)
		} else {
			logger.Debugf("As expected, current configuration has policy '%s'", policyName)
		}
	}
	if _, ok := pm.Manager([]string{policies.ApplicationPrefix}); ok {
//
		for _, policyName := range []string{
			policies.ChannelApplicationReaders,
			policies.ChannelApplicationWriters,
			policies.ChannelApplicationAdmins} {
			_, ok := pm.GetPolicy(policyName)
			if !ok {
				logger.Warningf("Current configuration has no policy '%s', this will likely cause problems in production systems", policyName)
			} else {
				logger.Debugf("As expected, current configuration has policy '%s'", policyName)
			}
		}
	}
	if _, ok := pm.Manager([]string{policies.OrdererPrefix}); ok {
		for _, policyName := range []string{policies.BlockValidation} {
			_, ok := pm.GetPolicy(policyName)
			if !ok {
				logger.Warningf("Current configuration has no policy '%s', this will likely cause problems in production systems", policyName)
			} else {
				logger.Debugf("As expected, current configuration has policy '%s'", policyName)
			}
		}
	}
}

