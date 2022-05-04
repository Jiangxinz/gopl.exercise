package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("%.2f", d) }

type database struct {
	data map[string]dollars
	sync.Mutex
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	defer db.Unlock()
	db.Lock()
	for item, price := range db.data {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	// /price?item=xxx
	item := req.URL.Query().Get("item")
	defer db.Unlock()
	db.Lock()
	price, ok := db.data[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func getItemAndPrice(req *http.Request) (string, float64, error) {
	req.ParseForm()
	item := req.FormValue("item")
	priceStr := req.FormValue("price")
	priceFloat, err := strconv.ParseFloat(priceStr, 32)
	return item, priceFloat, err
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	// /update?item=xxx&price=xxx
	item, price, err := getItemAndPrice(req)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "invalid price{%s}\n", req.FormValue("price"))
		return
	}
	defer db.Unlock()
	db.Lock()
	if _, ok := db.data[item]; !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item {%s} doesn't exist\n", item)
		return
	}
	db.data[item] = dollars(price)
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	// /add?item=xxx&price=xxx
	item, price, err := getItemAndPrice(req)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "invalid price{%s}\n", req.FormValue("price"))
		return
	}
	defer db.Unlock()
	db.Lock()
	if _, ok := db.data[item]; ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item {%s} already exists\n", item)
		return
	}
	db.data[item] = dollars(price)
}

func (db database) del(w http.ResponseWriter, req *http.Request) {
	// /del?item=xxx
	item := req.URL.Query().Get("item")
	defer db.Unlock()
	db.Lock()
	if _, ok := db.data[item]; !ok {
		msg := fmt.Sprintf("no such item: {%s}", item)
		http.Error(w, msg, http.StatusNotFound)
	}
	delete(db.data, item)
}

func main() {
	port := 8888
	fmt.Printf("Start HttpServer with Port %d\n", port)
	db := database{data: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/add", db.add)
	http.HandleFunc("/del", db.del)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}
