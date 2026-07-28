package main

import "net/http"

func (app *application) healthCheckHandler(wrt http.ResponseWriter, req *http.Request) {
	wrt.Write([]byte("ok"))
}
