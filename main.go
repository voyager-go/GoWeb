package main

import (
	"github.com/voyager-go/GoWeb/cmd"
	"github.com/voyager-go/GoWeb/pkg/logging"
	"os"
)

func main() {
	if err := cmd.App.Run(os.Args); err != nil {
		logging.Log.Error(err)
	}
}
