
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:52</date>
//</624455940376104960>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package common

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/hyperledger/fabric/cmd/common/comm"
	"github.com/hyperledger/fabric/cmd/common/signer"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	configFilePath := filepath.Join(os.TempDir(), fmt.Sprintf("config-%d.yaml", rand.Int()))
	fmt.Println(configFilePath)
	t.Run("save and load a config", func(t *testing.T) {
		c := Config{
			TLSConfig: comm.Config{
				CertPath:       "foo",
				KeyPath:        "foo",
				PeerCACertPath: "foo",
				Timeout:        time.Second * 3,
			},
			SignerConfig: signer.Config{
				KeyPath:      "foo",
				IdentityPath: "foo",
				MSPID:        "foo",
			},
		}

		err := c.ToFile(configFilePath)
		defer os.RemoveAll(configFilePath)
		assert.NoError(t, err)

		c2, err := ConfigFromFile(configFilePath)
		assert.NoError(t, err)
		assert.Equal(t, c, c2)
	})

	t.Run("bad config isn't saved", func(t *testing.T) {
		c := Config{}
		err := c.ToFile(configFilePath)
		assert.Contains(t, err.Error(), "config isn't valid")
	})

	t.Run("bad config isn't loaded", func(t *testing.T) {
		_, err := ConfigFromFile(filepath.Join("testdata", "not_a_yaml.yaml"))
		assert.Contains(t, err.Error(), "error unmarshaling YAML file")
	})

	t.Run("file that doesn't exist isn't loaded", func(t *testing.T) {
		_, err := ConfigFromFile(filepath.Join("testdata", "not_a_file.yaml"))
		assert.Contains(t, err.Error(), "no such file or directory")
	})
}

