package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	client, err := drive.NewService(ctx, option.WithCredentialsFile("tcm-dev-creds.json"))
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	resp, err := client.Files.Get("1_Q4yfxfbZtDdu_We8l4ci2_2COtnsFu0").Download()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	f, err := os.Create("downloaded.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	if _, err = io.Copy(f, resp.Body); err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("ok")
}
