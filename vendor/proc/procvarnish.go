package proc

import (
	"fmt"

	"procstrucs"

	"gopkg.in/h2non/gentleman.v2"
)

func VarnishPurge(Options procstrucs.OptionsPurge) {
	fmt.Println(Options)
	for i := 0; i < len(Options.HostGroup); i++ {
		fmt.Println(Options.HostGroup[i])
		cli := gentleman.New()
		cli.URL("http://ec2-54-173-70-41.compute-1.amazonaws.com:8080")
		req := cli.Request()
		req.Path("/" + Options.HostGroup[i])
		req.Method("PURGE")
		req.SetHeader("Host", "dev.capital.pe")
		req.SetHeader("X-Purge-Method", "regex")

		res, err := req.Send()
		if err != nil {
			fmt.Printf("Request error: %s\n", err)
			return
		}

		if !res.Ok {
			fmt.Printf("Invalid server response: %d\n", res.StatusCode)
			return
		}
		fmt.Sprintln("Request: %s", req.Body)
		fmt.Println("Body: %s", res.Header)
	}
}
