package secret

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEncryptCBCPKCS7(t *testing.T) {
	AuthorizationAesKey := "a5cc893b0db32f4f48963d516dff25cb"
	sendOrgMsg := []byte(`{"User":"system","Verify":"e10adc3949ba59abbe56e057f20f883e"}`)
	sendMsg, err := EncryptCBCPKCS7(AuthorizationAesKey, sendOrgMsg)
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(string(sendMsg))

	srcMsg, err := DecryptCBCPKCS7(AuthorizationAesKey, sendMsg)
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println(string(srcMsg))
}

func TestEncryptCBCPKCS5(t *testing.T) {
	AuthorizationAesKey := "a5cc893b0db32f4f48963d516dff25cb"
	sendOrgMsg := []byte(`{"User":"system","Verify":"e10adc3949ba59abbe56e057f20f883e"}`)
	sendMsg, err := EncryptCBCPKCS5(AuthorizationAesKey, sendOrgMsg, []byte(AuthorizationAesKey)[:16])
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(string(sendMsg))

	srcMsg, err := DecryptCBCPKCS5(AuthorizationAesKey, sendMsg, []byte(AuthorizationAesKey)[:16])
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println(string(srcMsg))
}

func TestEncryptCBCPKCS5tmp(t *testing.T) {
	AuthorizationAesKey := "a5cc893b0db32f4f48963d516dff25cb"
	sendOrgMsg := []byte(`{"User":"system","Verify":"e10adc3949ba59abbe56e057f20f883e"}`)
	sendMsg, err := EncryptCBCPKCS5(AuthorizationAesKey, sendOrgMsg, []byte(AuthorizationAesKey)[:16])
	if err != nil {
		logrus.Error(err)
		return
	}
	fmt.Println(string(sendMsg))

	srcMsg, err := DecryptCBCPKCS5(AuthorizationAesKey, sendMsg, []byte(AuthorizationAesKey)[:16])
	if err != nil {
		logrus.Error(err)
		return
	}

	fmt.Println(string(srcMsg))
}
