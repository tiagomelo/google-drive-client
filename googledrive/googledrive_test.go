// Copyright (c) 2023 Tiago Melo. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package googledrive

import (
	"context"
	"errors"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		name                string
		mockNewDriveService func(ctx context.Context, opts ...option.ClientOption) (*drive.Service, error)
		expectedError       error
	}{
		{
			name: "happy path",
			mockNewDriveService: func(ctx context.Context, opts ...option.ClientOption) (*drive.Service, error) {
				return new(drive.Service), nil
			},
		},
		{
			name: "error",
			mockNewDriveService: func(ctx context.Context, opts ...option.ClientOption) (*drive.Service, error) {
				return nil, errors.New("create service error")
			},
			expectedError: errors.New("creating drive service: create service error"),
		},
	}
	for _, tc := range testCases {
		newDriveService = tc.mockNewDriveService
		t.Run(tc.name, func(t *testing.T) {
			c, err := New(context.TODO(), "path/to/credsfile")
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf("expected error to be %v, got nil", tc.expectedError)
				}
				require.NotNil(t, c)
			}
		})
	}
}

func TestCreateFolder(t *testing.T) {
	testCases := []struct {
		name           string
		mockClosure    func(mfcc *mockFilesCreateCall, mfs *mockFileService, mds *mockDriveService)
		expectedOutput string
		expectedError  error
	}{
		{
			name: "happy path",
			mockClosure: func(mfcc *mockFilesCreateCall, mfs *mockFileService, mds *mockDriveService) {
				mfcc.file = &drive.File{
					Id: "someFolderId",
				}
			},
			expectedOutput: "someFolderId",
		},
		{
			name: "error",
			mockClosure: func(mfcc *mockFilesCreateCall, mfs *mockFileService, mds *mockDriveService) {
				mfcc.doErr = errors.New("create error")
			},
			expectedError: errors.New("creating folder someFolder under parent folders [parentFolder]: create error"),
		},
	}
	for _, tc := range testCases {
		mfcc := new(mockFilesCreateCall)
		mfs := new(mockFileService)
		mds := new(mockDriveService)
		mds.fs = mfs
		mfs.fcc = mfcc
		t.Run(tc.name, func(t *testing.T) {
			tc.mockClosure(mfcc, mfs, mds)
			client := &client{
				srv: mds,
			}
			output, err := client.CreateFolder("someFolder", "parentFolder")
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf("expected error to be %v, got nil", tc.expectedError)
				}
				require.Equal(t, tc.expectedOutput, output)
			}
		})
	}
}

func TestUploadFile(t *testing.T) {
	testCases := []struct {
		name           string
		file           *os.File
		mockClosure    func(mfcc *mockFilesCreateCall, mfs *mockFileService, mds *mockDriveService)
		expectedOutput string
		expectedError  error
	}{
		{
			name: "happy path",
			file: os.NewFile(555, "newFile"),
			mockClosure: func(mfcc *mockFilesCreateCall, mfs *mockFileService, mds *mockDriveService) {
				mfcc.file = &drive.File{
					Id: "someFileId",
				}
			},
			expectedOutput: "someFileId",
		},
		{
			name: "error",
			file: os.NewFile(555, "newFile"),
			mockClosure: func(mfcc *mockFilesCreateCall, mfs *mockFileService, mds *mockDriveService) {
				mfcc.doErr = errors.New("Do error")
			},
			expectedError: errors.New("creating file newFile under parent folders [parentFolder]: Do error"),
		},
	}
	for _, tc := range testCases {
		mfcc := new(mockFilesCreateCall)
		mfs := new(mockFileService)
		mds := new(mockDriveService)
		mds.fs = mfs
		mfs.fcc = mfcc
		mfcc.fcc = mfcc
		t.Run(tc.name, func(t *testing.T) {
			tc.mockClosure(mfcc, mfs, mds)
			client := &client{
				srv: mds,
			}
			output, err := client.UploadFile(tc.file, "parentFolder")
			if err != nil {
				if tc.expectedError == nil {
					t.Fatalf("expected no error, got %v", err)
				}
				require.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				if tc.expectedError != nil {
					t.Fatalf("expected error to be %v, got nil", tc.expectedError)
				}
				require.Equal(t, tc.expectedOutput, output)
			}
		})
	}
}

type mockFilesCreateCall struct {
	file  *drive.File
	doErr error
	fcc   filesCreateCall
}

func (m *mockFilesCreateCall) Do(opts ...googleapi.CallOption) (*drive.File, error) {
	return m.file, m.doErr
}

func (m *mockFilesCreateCall) Media(r io.Reader, options ...googleapi.MediaOption) filesCreateCall {
	return m.fcc
}

type mockFileService struct {
	fcc filesCreateCall
}

func (m *mockFileService) Create(file *drive.File) filesCreateCall {
	return m.fcc
}

type mockDriveService struct {
	fs fileService
}

func (m *mockDriveService) Files() fileService {
	return m.fs
}
