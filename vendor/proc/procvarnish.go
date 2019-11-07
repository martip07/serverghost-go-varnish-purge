package proc

import (
	"fmt"

	"procstrucs"

	"gopkg.in/h2non/gentleman.v2"
	"gopkg.in/h2non/gentleman.v2/plugins/headers"
)

func VarnishPurge(Options procstrucs.OptionsPurge) {
	fmt.Println(Options)
	for i, _ := range Options.PathGroup {
		fmt.Println("PATHS ===")
		fmt.Println(Options.PathGroup[i])
		for e, _ := range Options.VarnishGroup {
			fmt.Println("VARNISH ===")
			fmt.Println(Options.VarnishGroup[e])
			for h, _ := range Options.HostGroup {
				fmt.Println("HOST ===")
				fmt.Println(Options.HostGroup[h])
				cli := gentleman.New()
				cli.URL(Options.VarnishGroup[h])
				req := cli.Request()
				req.Path("/" + Options.PathGroup[h])
				req.Method("PURGE")
				req.SetHeader("Host", Options.HostGroup[h])
				req.Use(headers.Set("X-Purge-Method", "regex"))

				fmt.Printf("Request: %s", req)

				res, err := req.Send()
				if err != nil {
					fmt.Printf("Request error: %s\n", err)
					return
				}

				if !res.Ok {
					fmt.Printf("Invalid server response: %d\n", res.StatusCode)
					return
				}
				fmt.Printf("Response: %s", res.String())
				fmt.Println("Headers:", res.RawResponse)
			}
		}
	}
}
