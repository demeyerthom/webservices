# Webservices Client

### Usage:

    import (
        "fmt"
        webservices "github.com/demeyerthom/webservices/client"
        "github.com/hooklift/gowsdl/soap"
    )
    
    func main() {
        client := webservices.NewWebservicesClient(
            soap.NewClient("https://ws1.webservices.nl/soap"), 
            "username", 
            "secret",
         )
    
        response, err := client.MetaGetService(&webservices.MetaGetServiceRequestType{Service_name: "Meta"})
        if err != nil {
            panic(err)
        }
    
        fmt.Println(response.Out.Description)
    }

### Client (re)generation

Simply run `make`!
