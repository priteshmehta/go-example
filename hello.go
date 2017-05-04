/*
Go Example 
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"text/template"
	"github.com/ant0ine/go-json-rest/rest"
        "log"
        "net/http"
	"github.com/priteshtest/go-github/github"
)

func main() {
	fmt.Println("Hello World!");
        api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.SetApp(rest.AppSimple(func(w rest.ResponseWriter, r *rest.Request) {
		w.WriteJson(map[string]string{"Body": "Hello World!"})
	}))
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
	
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(ctx, "priteshtest", nil)
}
