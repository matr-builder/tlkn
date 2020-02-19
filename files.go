package tlkn

import (
	"context"
	"fmt"
	"io"
	"os"
)

// CopyFile ...
func CopyFile(ctx context.Context, src string, dest string) error {
	fromFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fromFile.Close()

	toFile, err := os.OpenFile(dest, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer toFile.Close()

	_, err = io.Copy(toFile, fromFile)
	if err != nil {
		return err
	}
	fmt.Printf("copied %s to %s", src, dest)
	return nil
}
