package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type OTXResult struct {
	HasNext    bool `json:"has_next"`
	ActualSize int  `json:"actual_size"`
	URLList    []struct {
		Domain   string `json:"domain"`
		URL      string `json:"url"`
		Hostname string `json:"hostname"`
		Httpcode int    `json:"httpcode"`
		PageNum  int    `json:"page_num"`
		FullSize int    `json:"full_size"`
		Paged    bool   `json:"paged"`
	} `json:"url_list"`
}

var (
	c = &http.Client{
		Timeout: time.Second * 15,
	}
	domains []string
	url     []string
)

func main() {
	flag.Parse()
	if flag.NArg() > 0 {
		domains = []string{flag.Arg(0)}
	} else {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			domains = append(domains, s.Text())
		}
	}
	for _, domain := range domains {
		page := 0
		for {
			r, err := c.Get(fmt.Sprintf("https://otx.alienvault.com/api/v1/indicators/hostname/%s/url_list?limit=50&page=%d", domain, page))
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
			defer r.Body.Close()
			bytes, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Fatal(err)
			}
			o := &OTXResult{}
			err = json.Unmarshal(bytes, o)
			if err != nil {
				log.Fatalf("Could not decode json: %s\n", err)
			}
			for _, url := range o.URLList {
				fmt.Println(url.URL)
			}
			if !o.HasNext {
				break
			}
			page++
		}
	}
}
