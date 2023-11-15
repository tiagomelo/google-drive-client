// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package googledrive

import (
	"io"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
)

// filesCreateCallWrapper is a wrapper around the Google Drive API's FilesCreateCall.
// It provides a convenient way to chain file creation and upload operations.
type filesCreateCallWrapper struct {
	fcc *drive.FilesCreateCall
}

// Do executes the file creation/upload call with the specified options.
// It delegates the call to the underlying Google Drive API's FilesCreateCall
// and returns the created file object or an error if the operation fails.
func (fccw *filesCreateCallWrapper) Do(opts ...googleapi.CallOption) (*drive.File, error) {
	return fccw.fcc.Do(opts...)
}

// Media sets the content of the file to be uploaded.
// It accepts an io.Reader for the file content and attaches it to the file creation call.
// Returns the filesCreateCallWrapper itself for further chaining of methods.
func (fccw *filesCreateCallWrapper) Media(r io.Reader, options ...googleapi.MediaOption) filesCreateCall {
	fccw.fcc.Media(r, options...)
	return fccw
}

// fileServiceWrapper wraps Google Drive's FilesService.
// It provides methods for file-related operations like file creation.
type fileServiceWrapper struct {
	srv *drive.FilesService
}

// Create initializes a new file creation call with the given file metadata.
// It returns a filesCreateCallWrapper for chaining file creation and upload configurations.
func (fsw *fileServiceWrapper) Create(file *drive.File) filesCreateCall {
	return &filesCreateCallWrapper{fsw.srv.Create(file)}
}

// driveServiceWrapper is a wrapper around Google Drive's Service.
// It provides access to the file operations through a fileService.
type driveServiceWrapper struct {
	srv *drive.Service
}

// Files returns a fileServiceWrapper that provides file-related operations.
func (dsw *driveServiceWrapper) Files() fileService {
	return &fileServiceWrapper{dsw.srv.Files}
}
