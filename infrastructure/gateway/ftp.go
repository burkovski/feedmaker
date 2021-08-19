package gateway

import (
	"context"
	"fmt"
	"io"
	"path"
	"time"

	"github.com/jlaffaye/ftp"
)

type (
	FtpConfig struct {
		Host        string
		Port        string
		ConnTimeout time.Duration `config:"conn_timeout"`
		Username    string
		Password    string
	}

	Dialer interface {
		DialTimeout(addr string, timeout time.Duration) (FtpConnection, error)
	}

	FtpConnection interface {
		Login(user, password string) error
		NoOp() error
		NameList(path string) ([]string, error)
		List(path string) ([]*ftp.Entry, error)
		Delete(path string) error
		RemoveDir(dir string) error
		MakeDir(path string) error
		Stor(path string, r io.Reader) error
		Quit() error
	}

	FtpGateway struct {
		Dialer Dialer
		Config FtpConfig
		conn   FtpConnection
	}
)

func (c FtpConfig) Addr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func (f *FtpGateway) Connect() error {
	conn, err := f.Dialer.DialTimeout(f.Config.Addr(), f.Config.ConnTimeout)
	if err != nil {
		return err
	}
	err = conn.Login(f.Config.Username, f.Config.Password)
	if err != nil {
		return err
	}
	f.conn = conn
	return nil
}

func (f *FtpGateway) reopenConnIfWasClosed() error {
	if f.connIsHealthy() {
		return nil
	}
	return f.Connect()
}

func (f *FtpGateway) connIsHealthy() bool {
	err := f.conn.NoOp()
	return err == nil
}

func (f *FtpGateway) DirExists(dir string) (bool, error) {
	if err := f.reopenConnIfWasClosed(); err != nil {
		return false, err
	}
	_, err := f.conn.NameList(dir)
	return err == nil, err
}

func (f *FtpGateway) CleanupDirContent(dir string) error {
	if err := f.reopenConnIfWasClosed(); err != nil {
		return err
	}
	return f.cleanupDirContentRecur(dir)
}

func (f *FtpGateway) cleanupDirContentRecur(dir string) error {
	entries, err := f.conn.List(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if isCurrentOrParentDir(entry) {
			continue
		}

		entryPath := path.Join(dir, entry.Name)

		if isNotDir(entry) {
			if err := f.conn.Delete(entryPath); err != nil {
				return err
			}
			continue
		}

		if err := f.cleanupDirContentRecur(entryPath); err != nil {
			return err
		}
		if err := f.conn.RemoveDir(entryPath); err != nil {
			return err
		}
	}

	return nil
}

func isCurrentOrParentDir(entry *ftp.Entry) bool {
	return entry.Name == "." || entry.Name == ".."
}

func isNotDir(entry *ftp.Entry) bool {
	return entry.Type != ftp.EntryTypeFolder
}

func (f *FtpGateway) MakeDir(dir string) error {
	if err := f.reopenConnIfWasClosed(); err != nil {
		return err
	}
	return f.conn.MakeDir(dir)
}

func (f *FtpGateway) Upload(ctx context.Context, path string, r io.Reader) error {
	if err := f.reopenConnIfWasClosed(); err != nil {
		return err
	}
	return f.conn.Stor(path, r)
}

func (f *FtpGateway) Disconnect() error {
	return f.conn.Quit()
}
