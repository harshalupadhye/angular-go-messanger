package main

import (
	"fmt"
	// "github.com/twilio/twilio-go"
	// openapi "github.com/twilio/twilio-go/rest/api/v2010"
	// "log"
	"net/smtp"
	"os"
)

func main() {
	// // //this is from twilio sms
	// accountSID := "AC8ef619f857b04845679cb6ee006a8ef7"
	// auth_token := "aff6cd9dc53d9bed2d45e4e792ad92c3"

	// // //create a server
	// client := twilio.NewRestClientWithParams(twilio.ClientParams{ //this is struct
	// 	Username: accountSID,
	// 	Password: auth_token,
	// })

	// // // params := &openapi.CreateMessageParams{} // this to create an API sms service
	// params := &openapi.CreateCallParams{} // this to create an API email service
	// params.SetTo("+917498171447")
	// params.SetFrom("+19378216738")
	// params.SetUrl("http://demo.twilio.com/docs/voice.xml")
	// // // params.SetBody("Hello from Harshal, This message was generated in order to test the application SMS notification functionality for Siemens Soris Adaptor")

	// // res, err := client.ApiV2010.CreateMessage(params)
	// res, err := client.ApiV2010.CreateCall(params)
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("message id", res.Sid)
	// }

	//email
	//configure sender
	from := "testinggolang1996@gmail.co"
	password := "Navkar@1996"

	//configure receiver
	toList := []string{"npupadhye64@gmail"}

	// host is address of server that the
	// sender's email address belongs,
	// in this case its gmail.
	// For e.g if your are using yahoo
	// mail change the address as smtp.mail.yahoo.com
	host := "smtp.gmail.com"

	// Its the default port of smtp server
	port := "587"

	msg := "Hello from Harshal, This message was generated in order to test the application email notification functionality for Siemens Soris Adaptor"

	 // We can't send strings directly in mail,
    // strings need to be converted into slice bytes
	body:= []byte(msg)

	// PlainAuth uses the given username and password to
    // authenticate to host and act as identity.
    // Usually identity should be the empty string,
    // to act as username.
	auth:= smtp.PlainAuth("", from, password, host)

	// SendMail uses TLS connection to send the mail
    // The email is sent to all address in the toList,
    // the body should be of type bytes, not strings
    // This returns error if any occured.
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)

	if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
 
    fmt.Println("Successfully sent mail to all user in toList")
}
