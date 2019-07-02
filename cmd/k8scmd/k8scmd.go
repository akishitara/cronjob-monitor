package main

import (
	"time"

	nexclient "github.com/akishitara/cronjob-monitor/pkg/client"
)

func main() {
	//debugger.YamlPrint(nexclient.ActiveCronjobsAllGet())
	nexclient.Run()
	time.Sleep(3600 * time.Second)
}
