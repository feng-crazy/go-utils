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
}
