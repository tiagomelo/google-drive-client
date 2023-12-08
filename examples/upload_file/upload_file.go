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
	FolderId      string `long:"folderId" description:"Folder Id" required:"true"`
	FilePath      string `long:"filePath" description:"File path" required:"true"`
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
	file, err := os.Open(opts.FilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	uploadedFileId, err := client.UploadFile(file, opts.FolderId)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("File created with ID: %s\n", uploadedFileId)
}
