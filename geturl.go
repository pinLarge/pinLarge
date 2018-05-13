package main

import (
    "log"
    "net/http"
	"io/ioutil"
	"strings"
	"os"
)

func main() {
	//listen on root path
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
    //grab the key from URL
	keys, ok := r.URL.Query()["key"]
    //return if keys are not available
    if !ok || len(keys) < 1 {
       return
    }
    // Query()["key"] will return an array of items, 
    // we only want the single item.
    key := keys[0]
	//return key
    log.Println("Url Param 'key' is: " + string(key))
	//key is suppossed to be a URL, calling the URL 
	resp, err := http.Get(key)
		if err != nil {
		panic(err)
	}
	//close the page
	defer resp.Body.Close()
	//read the response body in to feed
    feed, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//convert the feed in to string	
	feedx := BytesToString(feed)
	//get large images
	feedr := strings.Replace(feedx, "236x", "1200x", -1)
	//replace empty descriptions
	feeds := strings.Replace(feedr, "<description></description>", "<description>hid360.com</description>", -1)
	//replace all titles
	feedt := strings.Replace(feeds, "<title>", "<!-- title>", -1)
	feedu := strings.Replace(feedt, "</title>", "</title -->", -1)
	//preserve entry description
	feedi := strings.Replace(feedu, "<description>&lt;a", "<description1>&lt;a", -1)
	feedj := strings.Replace(feedi, "</description><pubDate>", "</description1><pubDate>", -1)
	//create new titles
	feedk := strings.Replace(feedj, "&gt;&lt;/a&gt;", "&gt;&lt;/a&gt;</description1><title>", -1)
	//feedl := strings.Replace(feedk, "</description1><pubDate>", "&lt;p&gt;&lt;a href=&quot;https://www.homeinteriordesign.org&quot;&gt;Here is an awesome home interior design and decor blog&lt;/a&gt;&lt;/p&gt;</title><pubDate>", -1)
	feedl := strings.Replace(feedk, "</description1><pubDate>", "&lt;p&gt;&lt;a href=&quot;http://www.homeinteriordesign.org/2018/02/short-guide-to-interior-decoration.html&quot;&gt;Short guide to interior decoration&lt;/a&gt;&lt;/p&gt;</title><pubDate>", -1)
	//restore description
	feedm := strings.Replace(feedl, "</description1>", "</description>", -1)
	feedn := strings.Replace(feedm, "<description1>", "<description>", -1)
	//add mandatory titles to channel
	feedo := strings.Replace(feedn, "<channel>", "<channel><title>Pinterest Large Image Feed by homeinteriordesign.org</title>", -1)
	//deal with stray p tags
	feedp := strings.Replace(feedo, "<title>&lt;/p&gt;", "<title>", -1)
	feedq := strings.Replace(feedp, "<description>&lt;p&gt;", "<description>", -1)
	//change back to byte array
	//fmt.Printf("%s\n", feedq)
	feedb := []byte(feedq)
	//return the response to the connections listening
	w.WriteHeader(resp.StatusCode)
	w.Write(feedb)	
}

func BytesToString(data []byte) string {
	return string(data[:])
}
