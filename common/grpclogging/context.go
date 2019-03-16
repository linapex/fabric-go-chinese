
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:39:56</date>
//</624455954699653120>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package grpclogging

import (
	"context"

	"go.uber.org/zap/zapcore"
)

type fieldKeyType struct{}

var fieldKey = &fieldKeyType{}

func ZapFields(ctx context.Context) []zapcore.Field {
	fields, ok := ctx.Value(fieldKey).([]zapcore.Field)
	if ok {
		return fields
	}
	return nil
}

func Fields(ctx context.Context) []interface{} {
	fields, ok := ctx.Value(fieldKey).([]zapcore.Field)
	if !ok {
		return nil
	}
	genericFields := make([]interface{}, len(fields))
	for i := range fields {
		genericFields[i] = fields[i]
	}
	return genericFields
}

func WithFields(ctx context.Context, fields []zapcore.Field) context.Context {
	return context.WithValue(ctx, fieldKey, fields)
}

