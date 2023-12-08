// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

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
	FolderName    string `short:"f" long:"folder" description:"Folder name" required:"true"`
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
	folderId, err := client.CreateFolder(opts.FolderName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("folder created with ID: %s\n", folderId)
}
