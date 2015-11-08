package main

import (
	"fmt"
	"github.com/everfore/gotest/mail"
	"github.com/shaalx/goutils"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/mailserv", mailService)
	http.ListenAndServe(":80", nil)
}

func index(rw http.ResponseWriter, req *http.Request) {
	rw.Write(goutils.ToByte("This is a webhooks service."))
}

func mailService(rw http.ResponseWriter, req *http.Request) {
	content := "Something changed:\n\n"
	reqb, err := ioutil.ReadAll(req.Body)
	if goutils.CheckErr(err) {
		content += err.Error() + "\n"
	}
	content += goutils.ToString(reqb)
	fmt.Println(req.RequestURI)
	fmt.Println(content)
	mail.SendMail(content)
}
