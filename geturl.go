package main

import (
    "log"
    "net/http"
	"io/ioutil"
	"strings"
	"os"
)

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

    keys, ok := r.URL.Query()["key"]
    
    if !ok || len(keys) < 1 {
        
       return
    }

    // Query()["key"] will return an array of items, 
    // we only want the single item.
    key := keys[0]
	//return key

    log.Println("Url Param 'key' is: " + string(key))
	resp, err := http.Get(key)
		if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
    feed, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	
    //fmt.Printf("%s\n", feed)
	
	feedx := BytesToString(feed)
	feedr := strings.Replace(feedx, "236x", "1200x", -1)
	
	//replace empty descriptions
	feeds := strings.Replace(feedr, "<description></description>", "<description>hid360.com</description>", -1)
    //replace empty titles
    feedt := strings.Replace(feeds, "<title></title>", "<title>More inspritations on hid360.com</title>", -1)
	//change to byte array
	//fmt.Printf("%s\n", feedt)
	feedb := []byte(feedq)
	
	w.WriteHeader(resp.StatusCode)
	w.Write(feedb)	
}

func BytesToString(data []byte) string {
	return string(data[:])
}