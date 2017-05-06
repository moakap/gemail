package gemail_test

import (
	"fmt"
	"github.com/moakap/gemail"
	"testing"
)

func TestSendEmail(t *testing.T) {
	//
	err := gemail.Send("user@example.com", "this is one test mail sent from go program", "this is the body.")
	if err != nil {
		fmt.Println("send mail failed: ", err.Error())
		t.Fail()
	}
}
