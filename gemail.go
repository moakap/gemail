package gemail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
)

// general account info, always get it from environment variables
const (
	servername = "smtp.exmail.qq.com:465" // port is needed
	username   = "274991933@qq.com"
	password   = "1314qqin"
)

// dial using TLS/SSL
func dial(addr string) (*tls.Conn, error) {
	/*
		// TLS config
		tlsconfig := &tls.Config{
			// InsecureSkipVerify controls whether a client verifies the
			// server's certificate chain and host name.
			// If InsecureSkipVerify is true, TLS accepts any certificate
			// presented by the server and any host name in that certificate.
			// In this mode, TLS is susceptible to man-in-the-middle attacks.
			// This should be used only for testing.
			InsecureSkipVerify: false,

			// ServerName indicates the name of the server requested by the client
			// in order to support virtual hosting. ServerName is only set if the
			// client is using SNI (see
			// http://tools.ietf.org/html/rfc4366#section-3.1).
			// ServerName: host,

			// MinVersion contains the minimum SSL/TLS version that is acceptable.
			// If zero, then TLS 1.0 is taken as the minimum.
			MinVersion: tls.VersionSSL30,

			// MaxVersion contains the maximum SSL/TLS version that is acceptable.
			// If zero, then the maximum version supported by this package is used,
			// which is currently TLS 1.2.
			MaxVersion: tls.VersionSSL30,
		}
	*/

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	return tls.Dial("tcp", addr, nil)
}

// compose message according to "from, to, subject, body"
func composeMsg(from string, to string, subject string, body string) (message string) {
	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from
	headers["To"] = to
	headers["Subject"] = subject

	// Setup message
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	return
}

// Send - send email over SSL
func Send(toAddr string, subject string, body string) (err error) {
	host, _, _ := net.SplitHostPort(servername)

	// get SSL connection
	conn, err := dial(servername)
	if err != nil {
		return
	}

	// create new SMTP client
	smtpClient, err := smtp.NewClient(conn, host)
	if err != nil {
		return
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", username, password, host)

	// auth the smtp client
	err = smtpClient.Auth(auth)
	if err != nil {
		return
	}

	// set To && From address, note that from address must be same as authorization user.
	from := mail.Address{"", username}
	to := mail.Address{"", toAddr}
	err = smtpClient.Mail(from.Address)
	if err != nil {
		return
	}

	err = smtpClient.Rcpt(to.Address)
	if err != nil {
		return
	}

	// Get the writer from SMTP client
	writer, err := smtpClient.Data()
	if err != nil {
		return
	}

	// compose message body
	message := composeMsg(from.String(), to.String(), subject, body)

	// write message to recp
	_, err = writer.Write([]byte(message))
	if err != nil {
		return
	}

	// close the writer
	err = writer.Close()
	if err != nil {
		return
	}

	// Quit sends the QUIT command and closes the connection to the server.
	smtpClient.Quit()

	return nil
}
