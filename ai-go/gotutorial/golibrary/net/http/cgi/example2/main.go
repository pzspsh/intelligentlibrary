/*
@File   : main.go
@Author : pan
@Time   : 2023-12-01 16:00:45
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cgi"
	"regexp"
)

type Handler struct{}

func (t Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	fmt.Println("In comming: " + request.Method)
	fmt.Println(request.URL)

	// user agent check
	userAgent := request.Header["User-Agent"][0]
	userAgentPattern, _ := regexp.Compile(`^(git)(.*)$`)
	userAgentSub := userAgentPattern.FindSubmatch([]byte(userAgent))

	if len(userAgentSub) == 0 {
		response.WriteHeader(405)
		return
	}
	// ignore auth
	// user, pass, ok := request.BasicAuth()

	// if !ok {
	//  response.Header().Set("WWW-Authenticate", `Basic realm="Git User Login"`)
	//  response.WriteHeader(401)
	//  return
	// }

	handler := new(cgi.Handler)
	handler.Path = "/usr/libexec/git-core/git-http-backend"
	pattern, _ := regexp.Compile(`^([a-zA-z0-9\/\.-]+\.git)(.*)$`)
	sub := pattern.FindSubmatch([]byte(request.URL.Path))

	// from user
	if len(sub) > 1 {
		handler.Env = append(handler.Env, "GIT_PROJECT_ROOT="+"/opt/gitrepo")
		handler.Env = append(handler.Env, "REMOTE_USER="+"dalong")
		handler.Env = append(handler.Env, "REMOTE_ADDR="+"localhost")
		handler.Env = append(handler.Env, "PGYER_UID="+"dalong")
	} else {
		response.WriteHeader(405)
		return
	}
	handler.Env = append(handler.Env, "GIT_HTTP_EXPORT_ALL=")
	request.Header.Del("REMOTE_USER")
	request.Header.Del("REMOTE_ADDR")
	request.Header.Del("PGYER-REPO")
	request.Header.Del("PGYER-REPO-USER")
	request.Header.Del("PGYER-REPO-ADDR")

	handler.ServeHTTP(response, request)
}

var RequestHandler Handler

func main() {

	http.HandleFunc("/", RequestHandler.ServeHTTP)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
