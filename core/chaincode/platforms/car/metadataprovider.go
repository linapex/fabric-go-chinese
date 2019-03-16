
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:04</date>
//</624455987532664832>

/*
版权所有，State Street Corp.保留所有权利。
γ
SPDX许可证标识符：Apache-2.0
**/

package car

import (
	"archive/tar"
	"bytes"
)

//元数据提供程序提供元数据
type MetadataProvider struct {
}

//getmetadataastarentries从chaincodedeploymentspec中提取metata数据
func (carMetadataProv *MetadataProvider) GetMetadataAsTarEntries() ([]byte, error) {
//这将转换汽车生成的元数据
//与生成的tar条目相同的tar条目
//其他平台。
//
//当前未实现，用户不能通过汽车指定元数据

	buf := bytes.NewBuffer(nil)
	tw := tar.NewWriter(buf)

	if err := tw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

