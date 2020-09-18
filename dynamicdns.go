package main

import (
	"fmt"
	"os"

	namecom "github.com/coolaj86/go-namedotcom/v4/namecom"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"git.ryanburnette.com/ryanburnette/dynamicdns/ipify"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	schedule := os.Getenv("SCHEDULE")
	if len(os.Getenv("SCHEDULE")) == 0 {
		schedule = "*/5 * * * *"
	}
	fmt.Println("Schedule:", schedule)

	fmt.Println("Notice: NAMECOM is the only supported API at the moment")
	api := os.Getenv("API")
	if len(api) == 0 {
		panic("API is required")
	}
	fmt.Printf("API: %s\n", api)

	nc := namecom.New(os.Getenv("NAMECOM_USERNAME"), os.Getenv("NAMECOM_API_TOKEN"))

	update(nc)

	c := cron.New()
	c.AddFunc(schedule, func() {
		update(nc)
	})
	c.Run()
}

func findRecord(nc *namecom.NameCom, hostname string) namecom.Record {
	resp, err := nc.ListRecords(&namecom.ListRecordsRequest{DomainName: hostname, PerPage: 999, Page: 1})
	if err != nil {
		panic(err)
	}

	var record namecom.Record
	for k := range resp.Records {
		v := resp.Records[k]
		if hostname+"." == v.Fqdn {
			record = *v
			break
		}
	}
	//fmt.Printf("%#v\n", record)

	return record
}

//type DNSAPI interface {}
//
//type UpdateRecordOptions struct {
//	API DNSAPI
//	TargetRecord string
//	NewIP net.IPAddr
//}

func update(nc *namecom.NameCom) {
	fmt.Println("Updating...")

	ip := ipify.GetIP()
	fmt.Println("IP:", ip)

	record := findRecord(nc, os.Getenv("DOMAIN"))

	if string(record.Answer) != string(ip) {
		record.Answer = ip
		_, err := nc.UpdateRecord(&record)
		if err != nil {
			panic(err)
		}
		fmt.Println("Record updated...")
	} else {
		fmt.Println("IP did not change...")
	}
}
