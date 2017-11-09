// dnsup.go
// Usage:  ./dnsup fqdn
package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/cloudflare/cloudflare-go"
)

// Get our IP address as a string via icanhazip.com
func getIP() string {
	resp, err := http.Get("http://icanhazip.com")
	if err != nil {
		log.Printf("Error reading IP from icanhazip.com, %v", err)
	}
	defer resp.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	return strings.Trim(buf.String(), "\n")

}

func main() {
	fqdn := os.Args[1]
	ip := getIP()

	api, err := cloudflare.New(os.Getenv("CF_API_KEY"), os.Getenv("CF_API_EMAIL"))
	if err != nil {
		log.Fatal(err)
	}

	fqdnS := strings.Split(fqdn, ".")
	zone := strings.Join(fqdnS[1:], ".")

	zoneID, err := api.ZoneIDByName(zone)

	dnsRec, err := api.DNSRecords(zoneID, cloudflare.DNSRecord{Name: fqdn})
	if err != nil {
		log.Fatal(err)
	}
	if dnsRec == nil {
		fmt.Println("FQDN not found, creating...")
		_, err = api.CreateDNSRecord(zoneID, cloudflare.DNSRecord{Name: fqdn, Type: "A", Content: ip})
		if err != nil {
			log.Fatalf("Record not created: %v", err)
		}
		fmt.Println("Done.")
	} else {
		fmt.Print("FQDN exists, updating... ")
		err = api.UpdateDNSRecord(zoneID, dnsRec[0].ID, cloudflare.DNSRecord{Name: fqdn, Content: ip})
		if err != nil {
			log.Fatalf("Record not updated: %v", err)
		}
		fmt.Println("Done.")
	}
}
