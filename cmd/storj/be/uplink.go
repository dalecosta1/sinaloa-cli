package be

import (
    "context"
    "io"
    "os"
    "storj.io/uplink"
)

type StorjService struct {
    Project *uplink.Project
}

func NewStorjService(secret string) (*StorjService, error) {
    ctx := context.Background()

    access, err := uplink.ParseAccess(secret)
    if err != nil {
        return nil, err
    }

    project, err := uplink.OpenProject(ctx, access)
    if err != nil {
        return nil, err
    }

    return &StorjService{Project: project}, nil
}

// AddFile uploads a file to the specified path in the Storj bucket
func (s *StorjService) AddFile(bucketName, pathStorj, pathFile string) error {
    ctx := context.Background()

    // Open the local file
    localFile, err := os.Open(pathFile)
    if err != nil {
        return err
    }
    defer localFile.Close()

    // Start an upload to the specified path in the Storj bucket
    upload, err := s.Project.UploadObject(ctx, bucketName, pathStorj, nil)
    if err != nil {
        return err
    }
    defer upload.Abort()

    // Copy the file data to the Storj upload object
    _, err = io.Copy(upload, localFile)
    if err != nil {
        return err
    }

    // Commit the upload to complete it
    return upload.Commit()
}

// AddFileBytes uploads a file,writing on file bytes, to the specified path in the Storj bucket
func (s *StorjService) AddFileBytes(bucketName, pathStorj string, content []byte) error {
    ctx := context.Background()

    // Start an upload to the specified path in the Storj bucket
    upload, err := s.Project.UploadObject(ctx, bucketName, pathStorj, nil)
    if err != nil {
        return err
    }
    defer upload.Abort()

    // Write the content to the Storj upload object
    _, err = upload.Write(content)
    if err != nil {
        return err
    }

    // Commit the upload to complete it
    return upload.Commit()
}

// GetFile downloads a file from the specified path in the Storj bucket
func (s *StorjService) GetFile(bucketName, pathStorj string) ([]byte, error) {
    ctx := context.Background()

    // Start a download from the specified path in the Storj bucket
    download, err := s.Project.DownloadObject(ctx, bucketName, pathStorj, nil)
    if err != nil {
        return nil, err
    }
    defer download.Close()

    // Read the file data from the Storj download object
    data, err := io.ReadAll(download)
    if err != nil {
        return nil, err
    }

    return data, nil
}

// DeleteFile deletes a file from the specified path in the Storj bucket
func (s *StorjService) DeleteFile(bucketName, pathStorj string) error {
    ctx := context.Background()

    // Delete the file from the specified path in the Storj bucket
    _, err := s.Project.DeleteObject(ctx, bucketName, pathStorj)
    if err != nil {
        return err
    }

    return nil
}