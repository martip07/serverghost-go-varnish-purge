package proc

import (
	"fmt"
	"os"

	"procstrucs"

	logger "github.com/izumin5210/gentleman-logger"
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
				// cli.URL(Options.VarnishGroup[h])
				// req := cli.Request()
				cli.Path(Options.PathGroup[i])
				// req.Path(Options.PathGroup[h])
				cli.Method("PURGE")
				// req.Method("PURGE")
				// cli.SetHeader("Host", Options.HostGroup[h])
				// req.SetHeader("Host", Options.HostGroup[h])
				cli.Use(headers.Set("Host", Options.HostGroup[h]))
				cli.Use(headers.Set("X-Purge-Method", "regex"))
				//req.Use(headers.Set("X-Purge-Method", "regex"))
				cli.Use(logger.New(os.Stdout))
				// req.Use(logger.New(os.Stdout))
				cli.Use(headers.Del("User-Agent"))

				res, err := cli.Request().BaseURL(Options.VarnishGroup[e]).Send()
				// res, err := req.Send()
				if err != nil {
					fmt.Printf("Request error: %s\n", err)
					return
				}

				if !res.Ok {
					fmt.Printf("Invalid server response: %d\n", res.StatusCode)
					return
				}
				fmt.Printf("Status: %d\n", res.StatusCode)
				fmt.Printf("Body: %s", res.String())
			}
		}
	}
}
