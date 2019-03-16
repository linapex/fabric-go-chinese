
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:25</date>
//</624456077441765376>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package nwo

import (
	"os"
	"os/exec"
)

type Command interface {
	Args() []string
	SessionName() string
}

type Enver interface {
	Env() []string
}

type WorkingDirer interface {
	WorkingDir() string
}

func ConnectsToOrderer(c Command) bool {
	for _, arg := range c.Args() {
		if arg == "--orderer" {
			return true
		}
	}
	return false
}

func NewCommand(path string, command Command) *exec.Cmd {
	cmd := exec.Command(path, command.Args()...)
	cmd.Env = os.Environ()
	if ce, ok := command.(Enver); ok {
		cmd.Env = append(cmd.Env, ce.Env()...)
	}
	if wd, ok := command.(WorkingDirer); ok {
		cmd.Dir = wd.WorkingDir()
	}
	return cmd
}

