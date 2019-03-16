
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:22</date>
//</624456065810960384>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package common

import (
	"sync/atomic"
)

//TLSCertificates聚合服务器和客户端TLS证书
type TLSCertificates struct {
TLSServerCert atomic.Value //*对等端的tls.certificate server证书
TLSClientCert atomic.Value //*对等端的tls.certificate客户端证书
}

