package main

import (
	"crypto/tls"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go domain1,domain2...")
		os.Exit(1)
	}
	domainList := strings.Split(os.Args[1], ",")
	today := time.Now()
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Domain", "Expires in (days)", ""})
	for _, domain := range domainList {
		func() {
			link := fmt.Sprintf("%s:443", domain)
			conn, err := tls.Dial("tcp", link, nil)
			if err != nil {
				t.AppendRow(table.Row{domain, -1, "--"})
			} else {
				defer conn.Close()
				certs := conn.ConnectionState().PeerCertificates
				for _, cert := range certs {
					expires := cert.NotAfter
					dateDiff := int32(expires.Sub(today).Hours() / 24)
					t.AppendRows([]table.Row{
						{domain, dateDiff, expires.Format("2006-01-02 15:04:05")},
					})
					break
				}
			}
		}()
	}
	t.SortBy([]table.SortBy{
		{Name: "Expires in (days)", Mode: table.AscNumeric},
	})
	fmt.Println(t.Render())
}
