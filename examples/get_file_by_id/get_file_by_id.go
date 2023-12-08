package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/tiagomelo/google-drive-client/googledrive"
)

var opts struct {
	CredsFilePath string `short:"c" long:"creds" description:"Creds file path" required:"true"`
	FileId        string `long:"fileId" description:"File Id" required:"true"`
}

func main() {
	if _, err := flags.ParseArgs(&opts, os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ctx := context.Background()
	client, err := googledrive.New(ctx, opts.CredsFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := client.GetFileById(opts.FileId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("file.Id: %v\n", file.Id)
}
