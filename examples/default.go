package main

import "github.com/overred/xout"

func main() {
	name := "XOut"

	// Redefine default global logger to enable debug levels.
	xout.Default = xout.NewPresetDefault(true)

	xout.Printf("Hello! This is the <fg=black;bg=cyan>%s</> Logger!\n", name)
	xout.Tracef("Hello! This is the <fg=black;bg=cyan>%s</> Logger!", name)
	xout.Debugf("Hello! This is the <fg=black;bg=cyan>%s</> Logger!", name)
	xout.Infof("Hello! This is the <fg=black;bg=cyan>%s</> Logger!", name)
	xout.Warnf("Hello! This is the <fg=black;bg=cyan>%s</> Logger!", name)
	xout.Errorf("Hello! This is the <fg=black;bg=cyan>%s</> Logger!", name)
}
