package main

import (
	"net/http"
	"encoding/json"
	"time"
	"fmt"
)

type Product struct {
	Title string
	Price string
}

func detailBasic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/template/detail_basic.html")
}

func detailAjax(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/template/detail_ajax.html")
}

//https://stackoverflow.com/questions/31622052/how-to-serve-up-a-json-response-using-go
func ajaxDetail(w http.ResponseWriter, r *http.Request) {

	product := Product{Title: "MAMA Jersey Top", Price: "$ 12.99"}
	j_product, err := json.Marshal(product)
	if err != nil {
		panic(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j_product)
}

func detailJson(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/template/detail_json.html")
}

func detailRegex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/template/detail_regex.html")
}

func detailHeader(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "assets/template/detail_header.html")
}

func ajaxdetailHeader(w http.ResponseWriter, r *http.Request) {
	// https://stackoverflow.com/questions/40096750/set-status-code-on-http-responsewriter
	// https://golang.org/src/net/http/status.go
	if isAjax(r) == false {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("403"))
		return
	}

	product := Product{Title: "MAMA Jersey Top", Price: "$ 12.99"}
	j_product, err := json.Marshal(product)
	if err != nil {
		panic(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j_product)

}

func detailCookie(w http.ResponseWriter, r *http.Request) {
	// https://astaxie.gitbooks.io/build-web-application-with-golang/content/zh/06.1.html

	expiration := time.Now().Add(time.Duration(30) * time.Hour)
	cookie := http.Cookie{Name: "token", Value: "233", Expires: expiration}
	http.SetCookie(w, &cookie)
	http.ServeFile(w, r, "assets/template/detail_cookie.html")
}

func ajaxdetailCookie(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("no cookie"))
		return
	}

	fmt.Println(cookie.Value)
	token := r.URL.Query().Get("token")
	fmt.Println(token)
	if token != cookie.Value {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("no cookie"))
		return
	}

	product := Product{Title: "MAMA Jersey Top", Price: "$ 12.99"}
	j_product, _ := json.Marshal(product)

	w.Header().Set("Content-Type", "application/json")
	w.Write(j_product)
}

func isAjax(r *http.Request) bool {
	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		return true
	}
	return false
}

func main() {
	http.HandleFunc("/detail_basic", detailBasic)
	http.HandleFunc("/detail_ajax", detailAjax)
	http.HandleFunc("/ajaxdetail", ajaxDetail)
	http.HandleFunc("/detail_json", detailJson)
	http.HandleFunc("/detail_regex", detailRegex)
	http.HandleFunc("/detail_header", detailHeader)
	http.HandleFunc("/ajaxdetail_header", ajaxdetailHeader)
	http.HandleFunc("/detail_cookie", detailCookie)
	http.HandleFunc("/ajaxdetail_cookie", ajaxdetailCookie)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.ListenAndServe(":8001", nil)
}
