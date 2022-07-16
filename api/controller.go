package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"code.cerinuts.io/uni/hypercocoagateway/shared"
)

type ApiController struct {
	Router *http.ServeMux
	Conf   *Config
}

// Route proxies requests restfully to the chain
func (ac *ApiController) Route() {
	ac.Router = http.NewServeMux()
	comparser := new(Comparser)
	comparser.Conf = ac.Conf
	comparser.Init()

	ac.Router.HandleFunc(shared.Hyperconfig.BaseURL+"/version", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Println("["+t.Format("2006-01-02 15:04:05")+"] Request", r.RequestURI)
		fmt.Fprintf(w, "%s\n", comparser.GetVersion())
	})

	ac.Router.HandleFunc(shared.Hyperconfig.BaseURL+"/assets", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			handle(w, r, func(w http.ResponseWriter, r *http.Request) string { return "" })
		}
		if r.Method == http.MethodGet {
			handle(w, r, comparser.GetAssets)
		}
		if r.Method == http.MethodPost {
			handle(w, r, comparser.AddAsset)
		}
	})

	ac.Router.HandleFunc(shared.Hyperconfig.BaseURL+"/assets/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handle(w, r, comparser.GetAsset)
		}
		if r.Method == http.MethodDelete {
			handle(w, r, comparser.DeleteAsset)
		}
		if r.Method == http.MethodPatch {
			handle(w, r, comparser.UpdateAsset)
		}
	})

}

// handle handles any requests according to the passed function
func handle(w http.ResponseWriter, r *http.Request, fn func(http.ResponseWriter, *http.Request) string) {
	t := time.Now()
	fmt.Println("["+t.Format("2006-01-02 15:04:05")+"] Request", r.RequestURI)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "HEAD, GET, POST, PUT, PATCH, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token")
	fmt.Fprintf(w, "%s", fn(w, r))
}

// Start the webserver. Note that Route must be called first
func (ac *ApiController) Start(host, port, sslport string) {
	if len(os.Args) > 1 && os.Args[1] == "--prod" {
		fmt.Println("Starting server listening to: " + host + ":" + port)
		fmt.Println("Starting SSL server listening to: " + host + ":" + sslport)
		go func() {
			log.Panic(http.ListenAndServeTLS(host+":"+sslport, ac.Conf.SslCert, ac.Conf.SslKey, ac.Router))
		}()
		log.Panic(http.ListenAndServe(host+":"+port, ac.Router))
	} else {
		fmt.Println("Starting server listening to: " + host + ":" + port)
		log.Panic(http.ListenAndServe(host+":"+port, ac.Router))
	}
}
