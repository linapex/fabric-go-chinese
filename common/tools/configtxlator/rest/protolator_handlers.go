
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:00</date>
//</624455970692534272>

/*
版权所有IBM Corp.2017保留所有权利。

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


package rest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/hyperledger/fabric/common/tools/protolator"
)

func getMsgType(r *http.Request) (proto.Message, error) {
	vars := mux.Vars(r)
msgName := vars["msgName"] //不到未设置

	msgType := proto.MessageType(msgName)
	if msgType == nil {
		return nil, fmt.Errorf("message name not found")
	}
	return reflect.New(msgType.Elem()).Interface().(proto.Message), nil
}

func Decode(w http.ResponseWriter, r *http.Request) {
	msg, err := getMsgType(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	err = proto.Unmarshal(buf, msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	var buffer bytes.Buffer
	err = protolator.DeepMarshalJSON(&buffer, msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	buffer.WriteTo(w)
}

func Encode(w http.ResponseWriter, r *http.Request) {
	msg, err := getMsgType(r)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, err)
		return
	}

	err = protolator.DeepUnmarshalJSON(r.Body, msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(data)
}

