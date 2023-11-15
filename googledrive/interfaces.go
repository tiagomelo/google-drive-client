// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package googledrive

import (
	"io"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// filesCreateCall abstracts the file creation and upload operations of Google Drive.
// It provides methods to set up the file metadata and content for upload.
type filesCreateCall interface {
	// Do executes the file creation/upload call with the specified options.
	// It returns the created file object or an error if the operation fails.
	Do(opts ...googleapi.CallOption) (*drive.File, error)

	// Media sets the content of the file to be uploaded.
	// It accepts an io.Reader for the file content and returns a filesCreateCall
	// for chaining further configurations.
	Media(r io.Reader, options ...googleapi.MediaOption) filesCreateCall
}

// fileService defines methods for file operations in Google Drive.
// It is a part of the driveService interface and primarily deals with file creation.
type fileService interface {
	// Create initializes a new file creation call with the given file metadata.
	// It returns a filesCreateCall to further configure and execute the file creation.
	Create(file *drive.File) filesCreateCall
}

// driveService abstracts the Google Drive service.
// It provides access to various functionalities like file operations through
// different sub-services.
type driveService interface {
	// Files returns a fileService that provides file-related operations
	// such as creating and uploading files.
	Files() fileService
}
