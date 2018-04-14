package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/julienschmidt/httprouter"
)

type TOML struct {
	Server Server `toml:"server"`
	APIs   []API  `toml:"api"`
}

type Server struct {
	Port string `toml:"port"`
}

type API struct {
	Method string `toml:"method"`
	URL    string `toml:"url"`
}

func main() {
	var cfg TOML
	if _, err := toml.DecodeFile("./config.toml", &cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Method	URL")

	router := httprouter.New()
	port := cfg.Server.Port
	for _, api := range cfg.APIs {
		m := strings.ToLower(api.Method)
		url := api.URL
		fp := "./resp" + url

		if _, err := os.Stat(fp); err != nil {
			// ファイルが存在しないので無視
			fmt.Println(fp + "が存在しません")
			continue
		}

		fmt.Printf("%s	http://localhost:%s%s\n", m, port, url)

		switch m {
		case "get":
			router.GET(url, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
				extIdx := strings.LastIndex(fp, ".")
				ext := fp[extIdx+1:]
				w.Header().Set("Content-Type", "application/"+ext+"; charset=UTF-8")
				b, err := ioutil.ReadFile(fp)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(w, string(b))
			})
		case "post":
			router.POST(url, func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
				extIdx := strings.LastIndex(fp, ".")
				ext := fp[extIdx+1:]
				w.Header().Set("Content-Type", "application/"+ext+"; charset=UTF-8")
				b, err := ioutil.ReadFile(fp)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(w, string(b))
			})
		default:
			fmt.Println(m + "は不正なリクエストメソッドです。")
		}
	}

	log.Fatal(http.ListenAndServe(":"+port, router))
}
