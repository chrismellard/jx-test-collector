package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sethvargo/go-envconfig/pkg/envconfig"

	"github.com/jenkins-x/jx-test-collector/pkg/tailer"
)

func main() {
	ctx := context.Background()

	o := &tailer.Options{}
	if err := envconfig.Process(ctx, o); err != nil {
		log.Fatal(err)
	}

	if err := o.Run(); err != nil {
		_, err = fmt.Fprintf(os.Stderr, "ERROR: %s\n", err.Error())
		if err != nil {
			os.Exit(2)
		}
		os.Exit(1)
	}
}
