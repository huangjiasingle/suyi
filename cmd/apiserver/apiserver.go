package main

import (
	"fmt"
	"os"

	"github.com/huangjiasingle/suyi/cmd/apiserver/app"
)

func main() {
	cmd := app.NewAPIServerCommand()
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
