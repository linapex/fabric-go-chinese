
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:51</date>
//</624455933828796416>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
*/

package bccsp

import (
	"fmt"
)

type IdemixIIssuerPublicKeyImporterErrorType int

const (
	IdemixIssuerPublicKeyImporterUnmarshallingError IdemixIIssuerPublicKeyImporterErrorType = iota
	IdemixIssuerPublicKeyImporterHashError
	IdemixIssuerPublicKeyImporterValidationError
	IdemixIssuerPublicKeyImporterNumAttributesError
	IdemixIssuerPublicKeyImporterAttributeNameError
)

type IdemixIssuerPublicKeyImporterError struct {
	Type     IdemixIIssuerPublicKeyImporterErrorType
	ErrorMsg string
	Cause    error
}

func (r *IdemixIssuerPublicKeyImporterError) Error() string {
	if r.Cause != nil {
		return fmt.Sprintf("%s: %s", r.ErrorMsg, r.Cause)
	}

	return fmt.Sprintf("%s", r.ErrorMsg)
}

