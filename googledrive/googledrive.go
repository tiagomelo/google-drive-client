// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package googledrive

import (
	"context"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

// For ease of unit testing.
var newDriveService = drive.NewService

// Client defines the interface for interacting with Google Drive.
// It abstracts the functionality for creating folders and uploading files.
type Client interface {
	// CreateFolder creates a new folder in Google Drive.
	// folderName is the name of the new folder. parentFolders is an optional list of
	// parent folder IDs where the new folder will be created.
	// Returns the ID of the created folder or an error if the operation fails.
	CreateFolder(folderName string, parentFolders ...string) (string, error)

	// UploadFile uploads a file to Google Drive.
	// file is a pointer to the os.File object to be uploaded. parentFolders is an
	// optional list of parent folder IDs where the file will be uploaded.
	// Returns the ID of the uploaded file or an error if the operation fails.
	UploadFile(file *os.File, parentFolders ...string) (string, error)
}

// client is an implementation of the Client interface.
// It provides methods to interact with Google Drive.
type client struct {
	srv driveService
}

func (c *client) CreateFolder(folderName string, parentFolders ...string) (string, error) {
	const mimeType = "application/vnd.google-apps.folder"
	fileSrv := c.srv.Files()
	call := fileSrv.Create(&drive.File{
		Name:     folderName,
		MimeType: mimeType,
		Parents:  parentFolders,
	})
	createdFolder, err := call.Do()
	if err != nil {
		return "", errors.Wrapf(err, "creating folder %s under parent folders %v", folderName, parentFolders)
	}
	return createdFolder.Id, nil
}

func (c *client) UploadFile(file *os.File, parentFolders ...string) (string, error) {
	driveFile := &drive.File{
		Name:    filepath.Base(file.Name()),
		Parents: parentFolders,
	}
	fileSrv := c.srv.Files()
	call := fileSrv.Create(driveFile).Media(file)
	createdFile, err := call.Do()
	if err != nil {
		return "", errors.Wrapf(err, "creating file %s under parent folders %v", file.Name(), parentFolders)
	}
	return createdFile.Id, nil
}

// New creates a new instance of a Client.
// ctx is the context for the drive service, and credsFilePath is the path to the
// credentials file for Google Drive API.
// Returns an instance of Client or an error if the drive service cannot be created.
func New(ctx context.Context, credsFilePath string) (Client, error) {
	srv, err := newDriveService(ctx, option.WithCredentialsFile(credsFilePath))
	if err != nil {
		return nil, errors.Wrap(err, "creating drive service")
	}
	return &client{
		srv: &driveServiceWrapper{srv},
	}, nil
}
