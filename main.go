package main

import (
	"flag"
	"fmt"
	purgeOptions "procstrucs"
	purgeProc "procvarnish"
)

func main() {
	fmt.Println("Howdy serverghost go varnish purge")

	var flagHost purgeOptions.ArrayHosts
	var flagVarnish purgeOptions.ArrayVarnishUris
	var flagPath purgeOptions.ArrayPathUris

	flag.Var(&flagHost, "hostlist", "List of your apps hostname, example: domain.com")
	flag.Var(&flagVarnish, "varnishlist", "List of your varnish servers, example: varnish1.com or 1.2.3.4")
	flag.Var(&flagPath, "pathlist", "List of you apps uri paths, example: /home or /dashboard")

	flag.Parse()

	var Options purgeOptions.OptionsPurge
	Options = purgeOptions.OptionsPurge{
		HostGroup:    flagHost,
		VarnishGroup: flagVarnish,
		PathGroup:    flagPath,
	}
	purgeProc.VarnishPurge(Options)
}
