package client

import (
	debugger "github.com/akishitara/cronjob-monitor/pkg/debugger"
)

// Run Debug用
func Run() {
	debugger.YamlPrint(ActiveCronjobsAllGet())
}
