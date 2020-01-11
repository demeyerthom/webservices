package client

import "github.com/hooklift/gowsdl/soap"

//TODO: fix generation to add these automatically
type (
	//GYearMonth
	GYearMonth string

	//GYear
	GYear string
)

//WebservicesClient provides an interface wrapper around the generated code
type WebservicesClient interface {
	WebservicesNLPortType
}

//NewWebservicesClient creates a new WebservicesClient
func NewWebservicesClient(soapClient *soap.Client, username string, password string) WebservicesClient {
	soapClient.AddHeader(HeaderLoginType{
		Username: username,
		Password: password,
	})

	generated := NewWebservicesNLPortType(soapClient)

	return generated
}
