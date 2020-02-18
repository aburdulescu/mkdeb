package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func CopyFile(src, dst string, mode os.FileMode) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	if err := out.Sync(); err != nil {
		return err
	}
	if err := out.Chmod(mode); err != nil {
		return err
	}
	return nil
}

func CopyDir(src string, dst string) error {
	if err := os.MkdirAll(dst, os.ModePerm); err != nil {
		return err
	}
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			if err := CopyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			if err := CopyFile(srcPath, dstPath, entry.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}
