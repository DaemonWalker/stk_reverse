package main

import (
	"errors"
	"github.com/fatih/color"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func main() {
	passUri, err := getPassUri()
	if err != nil {
		color.Red("error: pass uri is illegal reason: %s", err.Error())
		return
	}

	proxyUri, err := getProxyUri()
	if err != nil {
		color.Yellow("warning: proxy uri is not set, reason: %s", err.Error())
	}

	listen := getParameter("--listen", "STK_REVERSE_LISTEN")
	if len(listen) == 0 {
		listen = ":8080"
	}

	proxy := httputil.NewSingleHostReverseProxy(passUri)
	proxy.ErrorLog = log.Default()
	baseDirector := proxy.Director
	proxy.Director = func(request *http.Request) {
		baseDirector(request)
		request.Host = passUri.Host
	}

	if proxyUri != nil {
		proxy.Transport = &http.Transport{Proxy: http.ProxyURL(proxyUri)}
	}

	http.HandleFunc("/", proxy.ServeHTTP)
	color.Green("info: try to start server on %s\n", listen)
	color.Set(color.FgYellow)
	color.Red("error: server start failed, reason: %s\n", http.ListenAndServe(listen, nil))
}

func getPassUri() (*url.URL, error) {
	value := getParameter("--pass", "STK_REVERSE_PASS")
	if len(value) == 0 {
		return nil, errors.New("parameter PASS is empty, consider using --pass in command line or setting environment value STK_REVERSE_PASS")
	}
	return url.Parse(value)
}

func getProxyUri() (*url.URL, error) {
	value := getParameter("--proxy", "STK_REVERSE_PROXY")
	if len(value) == 0 {
		return nil, errors.New("parameter PROXY is empty, consider using --proxy in command line or setting environment value STK_REVERSE_PROXY")
	}
	return url.Parse(value)
}

func getParameter(cmdName string, envName string) string {
	args := os.Args
	for i := 0; i < len(args); i++ {
		if args[i] == cmdName {
			return args[i+1]
		}
	}
	return os.Getenv(envName)
}
