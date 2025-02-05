package backup

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// hash will be made up to filename and pth, so if file is renamed hash will be different or any change
func DirHash(path string) (string, error) {
	hash := md5.New()

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		io.WriteString(hash, path)
		fmt.Fprintf(hash, "%v", info.IsDir())
		fmt.Fprintf(hash, "%v", info.ModTime())
		fmt.Fprintf(hash, "%v", info.Mode())
		fmt.Fprintf(hash, "%v", info.Name())
		fmt.Fprintf(hash, "%v", info.Size())
		return nil

	})

	if err != nil {
		return "", err
	}
	//Sum method calculates final hash value with specified values appended
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
