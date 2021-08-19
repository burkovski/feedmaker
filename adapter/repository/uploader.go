package repository

import (
	"context"
	"fmt"
	"io"
	"path"

	"github.com/rs/zerolog/log"
)

const (
	rootDir = "/"
)

type (
	FtpGateway interface {
		DirExists(dir string) (bool, error)
		CleanupDirContent(dir string) error
		MakeDir(path string) error
		Upload(ctx context.Context, path string, r io.Reader) error
	}

	ftpUploader struct {
		ftp              FtpGateway
		generationType   string
		inStream         <-chan io.ReadCloser
		uploadedFilesNum uint
		onUpload         func(uploadedFilesNum uint)
	}
)

func NewFtpUploader(ftpGateway FtpGateway, generationType string, inStream <-chan io.ReadCloser) *ftpUploader {
	return &ftpUploader{
		ftp:            ftpGateway,
		generationType: generationType,
		inStream:       inStream,
		onUpload:       func(uploadedFilesNum uint) {},
	}
}

func (u *ftpUploader) UploadFiles(ctx context.Context) error {
	uploadDir := path.Join(rootDir, u.generationType)

	if err := u.prepareUploadDir(uploadDir); err != nil {
		return err
	}

	for {
		select {
		case file, isOpen := <-u.inStream:
			if !isOpen {
				return nil
			}

			filename := fmt.Sprintf("%s_%d.csv", u.generationType, u.uploadedFilesNum)
			filepath := path.Join(rootDir, u.generationType, filename)

			if err := u.ftp.Upload(ctx, filepath, file); err != nil {
				return err
			}

			u.uploadedFilesNum++
			u.onUpload(u.uploadedFilesNum)
			if err := file.Close(); err != nil {
				log.Error().Err(err).Msgf("Cannot close file after uploading")
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (u *ftpUploader) prepareUploadDir(uploadDir string) error {
	if exists, err := u.ftp.DirExists(uploadDir); err != nil {
		return err
	} else if exists {
		return u.ftp.CleanupDirContent(uploadDir)
	}

	return u.ftp.MakeDir(uploadDir)
}

func (u *ftpUploader) OnUpload(callback func(uploadedFilesNum uint)) {
	u.onUpload = callback
}
