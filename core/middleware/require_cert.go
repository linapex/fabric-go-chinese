
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:15</date>
//</624456036777988096>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package middleware

import (
	"net/http"
)

type requireCert struct {
	next http.Handler
}

//RequireCert用于确保验证的TLS客户端证书
//用于身份验证。
func RequireCert() Middleware {
	return func(next http.Handler) http.Handler {
		return &requireCert{next: next}
	}
}

func (r *requireCert) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch {
	case req.TLS == nil:
		fallthrough
	case len(req.TLS.VerifiedChains) == 0:
		fallthrough
	case len(req.TLS.VerifiedChains[0]) == 0:
		w.WriteHeader(http.StatusUnauthorized)
	default:
		r.next.ServeHTTP(w, req)
	}
}

