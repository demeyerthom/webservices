.DEFAULT: client
.PHONY: client client/gowsdl client/generate client/clean

## Work with client

client: client/webservices.wsdl client/gowsdl client/generate client/clean

client/clean:
	@rm ./client/webservices.wsdl

client/webservices.wsdl:
	@curl https://api.webservices.nl/soap_doclit?wsdl > ./client/tmp_webservices.wsdl
	@sed 's/Webservices.nl/WebservicesNL/g' ./client/tmp_webservices.wsdl > ./client/webservices.wsdl
	@rm ./client/tmp_webservices.wsdl

client/gowsdl:
	@go get github.com/hooklift/gowsdl/...

client/generate:
	@gowsdl -o generated.go -p client -make-public false ./client/webservices.wsdl

