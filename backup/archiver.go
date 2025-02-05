package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

type Archiver interface {
	Archive(src, dest string) error //this method takes the source and dest. paths and returns an error
}

type zipper struct{}

var ZIP Archiver = (*zipper)(nil)

func (z *zipper) Archive(src, dest string) error {
	// ensures that dest. directory exists
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err	
	}

	// creatin new file specified by dest path
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()

	// will write to file we just crreated and defer closing of writer
	w := zip.NewWriter(out)
	defer w.Close()

			// iterates over source directory
												//callback function to be callued for every item
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}

		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()

		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, in)
		if err != nil {
			return err
		}

		return nil
	})

}
