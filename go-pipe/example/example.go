package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/yongman/go/go-pipe"
)

func main() {
	var b bytes.Buffer
	if err := pipe.Command(&b,
		exec.Command("ls", "/-l"),
		exec.Command("grep", "hero"),
	); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(os.Stdout, &b); err != nil {
		log.Fatal(err)
	}
}
