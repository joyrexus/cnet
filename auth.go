package cnet

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/ldap.v1"
)

var (
	server string = "ldap.uchicago.edu"
	port   uint16 = 389
	baseDN string = "dc=uchicago,dc=edu"
	user   string = "uid=%s,ou=people,dc=uchicago,dc=edu"
)

func Authenticate(name, passwd string) error {
	address := fmt.Sprintf("%s:%d", server, port)
	l, err := ldap.Dial("tcp", address)
	if err != nil {
		return fmt.Errorf("Unable to connect to %s: %s\n", address, err)
	}
	defer l.Close()
	l.Debug = false

	err = l.StartTLS(&tls.Config{ServerName: server})
	if err != nil {
		return fmt.Errorf("Unable to secure connection to %s: %s\n", server, err)
	}

	entry := fmt.Sprintf(user, name)
	err = l.Bind(entry, passwd)
	if err != nil {
		return fmt.Errorf("Unable to authenticate %s: %s\n", name, err)
	}

	return nil // success!
}
