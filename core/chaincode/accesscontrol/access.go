
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:40:01</date>
//</624455977889959936>

/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package accesscontrol

import (
	"errors"
	"fmt"

	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/common/crypto/tlsgen"
	"github.com/hyperledger/fabric/common/flogging"
	pb "github.com/hyperledger/fabric/protos/peer"
	"google.golang.org/grpc"
)

var logger = flogging.MustGetLogger("chaincode.accesscontrol")

//certandprivkeypair包含一个证书
//及其相应的base64格式的私钥
type CertAndPrivKeyPair struct {
//
	Cert string
//密钥-对应证书的私钥
	Key string
}

type Authenticator struct {
	mapper *certMapper
}

func (auth *Authenticator) Wrap(srv pb.ChaincodeSupportServer) pb.ChaincodeSupportServer {
	return newInterceptor(srv, auth.authenticate)
}

//new authenticator返回可以包装链码服务的新authenticator
func NewAuthenticator(ca tlsgen.CA) *Authenticator {
	return &Authenticator{
		mapper: newCertMapper(ca.NewClientCertKeyPair),
	}
}

//generate返回一对证书和私钥，
//并将证书哈希与给定的
//链码名称
func (ac *Authenticator) Generate(ccName string) (*CertAndPrivKeyPair, error) {
	cert, err := ac.mapper.genCert(ccName)
	if err != nil {
		return nil, err
	}
	return &CertAndPrivKeyPair{
		Key:  cert.PrivKeyString(),
		Cert: cert.PubKeyString(),
	}, nil
}

func (ac *Authenticator) authenticate(msg *pb.ChaincodeMessage, stream grpc.ServerStream) error {
	if msg.Type != pb.ChaincodeMessage_REGISTER {
		logger.Warning("Got message", msg, "but expected a ChaincodeMessage_REGISTER message")
		return errors.New("First message needs to be a register")
	}

	chaincodeID := &pb.ChaincodeID{}
	err := proto.Unmarshal(msg.Payload, chaincodeID)
	if err != nil {
		logger.Warning("Failed unmarshaling message:", err)
		return err
	}
	ccName := chaincodeID.Name
//从流获取证书
	hash := extractCertificateHashFromContext(stream.Context())
	if len(hash) == 0 {
		errMsg := fmt.Sprintf("TLS is active but chaincode %s didn't send certificate", ccName)
		logger.Warning(errMsg)
		return errors.New(errMsg)
	}
//在地图绘制器中查找
	registeredName := ac.mapper.lookup(certHash(hash))
	if registeredName == "" {
		errMsg := fmt.Sprintf("Chaincode %s with given certificate hash %v not found in registry", ccName, hash)
		logger.Warning(errMsg)
		return errors.New(errMsg)
	}
	if registeredName != ccName {
		errMsg := fmt.Sprintf("Chaincode %s with given certificate hash %v belongs to a different chaincode", ccName, hash)
		logger.Warning(errMsg)
		return fmt.Errorf(errMsg)
	}

	logger.Debug("Chaincode", ccName, "'s authentication is authorized")
	return nil
}

