package main

import (
	"io"
	"os"
)

func cp(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	// create destination, requiring that it not exist
	out, err := os.OpenFile(dst, os.O_CREATE|os.O_RDWR|os.O_EXCL, 0666)
	if os.IsExist(err) {
		// assume the file is correct, and we don't need to copy
		// this is necessary to make the init container idempotent,
		// and avoid crashes if it is somehow restarted
		return nil
	} else if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	// make the file executable
	if err = out.Chmod(os.ModePerm); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := cp("/app/worker", "/pach-bin/worker"); err != nil {
		panic(err)
	}
	if err := cp("/app/pachctl", "/pach-bin/pachctl"); err != nil {
		panic(err)
	}
}
