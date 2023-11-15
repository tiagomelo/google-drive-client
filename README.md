# google-drive-client

A simple Go client for interacting with Google Drive. This client provides an easy way to create folders and upload files to Google Drive using the Google Drive API.

More functionalities will be added.

## setup

Before using this client, you need to set up a Google Cloud Project and enable the Google Drive API. Follow these steps:

1. Create a new project in the [Google Cloud Console](https://console.cloud.google.com/).
2. Enable the [Google Drive API](https://console.cloud.google.com/apis/library/drive.googleapis.com) for your project.
3. Create credentials (service account key) for your project:
   - Go to the "Credentials" page.
   - Click "Create Credentials" and select "Service account key".
   - Choose or create a new service account.
   - Set the role to `Editor` or another role that grants the necessary permissions.
   - Select JSON for the key type and download the file.
4. Save the downloaded JSON file in the `creds` directory.

For more detailed instructions, see the [Google Drive API documentation](https://developers.google.com/drive/api/v3/quickstart/go).

## example

Here's a basic example of how to use this client:

```
package main

import (
    "context"
    "fmt"
    "os"

    "github.com/tiagomelo/google-drive-client/googledrive"
)

func main() {
    ctx := context.Background()
    client, err := googledrive.New(ctx, credsFilePath)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    const parentFolderId = `1LTtYxqYUMoA1IzHmlx_RMVdkWgtskMxO`
    folder, err := client.CreateFolder("test folder", parentFolderId)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("folder.Id: %v\n", folder)

    file, err := os.Open("test.txt")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer file.Close()

    uploadedFile, err := client.UploadFile(file, folder)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Printf("uploadedFile: %v\n", uploadedFile)
}
```

## unit tests

```
make test
```