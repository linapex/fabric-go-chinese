
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:04</date>
//</624455988497354752>

/*
版权所有2017-greg haskins<gregory.haskins@gmail.com>

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


package golang

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"
)

//runprogram non nil env、超时（通常为秒或毫秒）、程序名和参数
func runProgram(env Env, timeout time.Duration, pgm string, args ...string) ([]byte, error) {
	if env == nil {
		return nil, fmt.Errorf("<%s, %v>: nil env provided", pgm, args)
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	cmd := exec.CommandContext(ctx, pgm, args...)
	cmd.Env = flattenEnv(env)
	stdErr := &bytes.Buffer{}
	cmd.Stderr = stdErr

	out, err := cmd.Output()

	if ctx.Err() == context.DeadlineExceeded {
		err = fmt.Errorf("timed out after %s", timeout)
	}

	if err != nil {
		return nil,
			fmt.Errorf(
				"command <%s %s>: failed with error: \"%s\"\n%s",
				pgm,
				strings.Join(args, " "),
				err,
				string(stdErr.Bytes()))
	}
	return out, nil
}

//逻辑灵感来源：https://dave.cheney.net/2014/09/14/go-list-your-swiss-army-nike
func list(env Env, template, pkg string) ([]string, error) {
	if env == nil {
		env = getEnv()
	}

	lst, err := runProgram(env, 60*time.Second, "go", "list", "-f", template, pkg)
	if err != nil {
		return nil, err
	}

	return strings.Split(strings.Trim(string(lst), "\n"), "\n"), nil
}

func listDeps(env Env, pkg string) ([]string, error) {
	return list(env, "{{ join .Deps \"\\n\"}}", pkg)
}

func listImports(env Env, pkg string) ([]string, error) {
	return list(env, "{{ join .Imports \"\\n\"}}", pkg)
}

