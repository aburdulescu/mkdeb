package main

import (
	"os"
	"path/filepath"
)

func makeOutDir(outputDir string, p Package) error {
	if err := mkControl(outputDir, p.Control); err != nil {
		return err
	}
	if err := mkArtifacts(outputDir, p.Artifacts); err != nil {
		return err
	}
	if err := mkScripts(outputDir, p.Scripts); err != nil {
		return err
	}
	return nil
}

func mkControl(outputDir string, c map[string]string) error {
	debianDirPath := filepath.Join(outputDir, "DEBIAN")
	if err := os.MkdirAll(debianDirPath, os.ModePerm); err != nil {
		return err
	}
	control, err := os.Create(filepath.Join(debianDirPath, "control"))
	if err != nil {
		return err
	}
	defer control.Close()
	for k, v := range c {
		if _, err := control.WriteString(k + ": " + v + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func mkArtifacts(outDirPath string, a map[string]string) error {
	for src, dst := range a {
		srcPath := filepath.Clean(src)
		fi, err := os.Stat(srcPath)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(outDirPath, dst)
		if fi.IsDir() {
			if CopyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(dstPath), os.ModePerm); err != nil {
				return err
			}
			if CopyFile(srcPath, dstPath, fi.Mode()); err != nil {
				return err
			}
		}
	}
	return nil
}

func mkScripts(outDirPath string, s map[string]string) error {
	debianDirPath := filepath.Join(outDirPath, "DEBIAN")
	for name, content := range s {
		path := filepath.Join(debianDirPath, name)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		_, err = f.WriteString("#!/bin/bash\n" + content + "\n")
		if err != nil {
			return err
		}
		if err := f.Chmod(0755); err != nil {
			return err
		}
		f.Close()
	}
	return nil
}
