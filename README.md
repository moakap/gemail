# gemail
Send email over SMTP by golang. 

## Example

mail.go
```
package main

import (
	"fmt"
	"github.com/moakap/gemail"
)

func main() {
  // send email from go program.
  err := gemail.Send("user@example.com", "this is one test mail sent from go program", "this is the body.")
	if err != nil {
		fmt.Println("send mail failed: ", err.Error())
	}
}

```
