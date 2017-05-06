package gemail_test

import (
	"fmt"
	"git.coding.net/moakap/gemail"
	"testing"
)

func TestSendEmail(t *testing.T) {
	//
	err := email.Send("jun.liu@obenben.com", "this is one test mail sent from go program", "this is the body.")
	if err != nil {
		fmt.Println("send mail failed: ", err.Error())
		t.Fail()
	}
}
