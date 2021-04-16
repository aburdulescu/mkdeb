package main

import (
	"errors"
	"os"
	"path/filepath"
)

type Metadata struct {
	Control map[string]string `yaml:"control"`
	Scripts map[string]string `yaml:"scripts"`
	Data    map[string]string `yaml:"data"`
}

func (m Metadata) Validate() error {
	if err := m.validateControl(); err != nil {
		return err
	}
	if m.Data == nil {
		return errors.New("data is empty or missing")
	}
	if err := m.validateData(); err != nil {
		return err
	}
	if m.Scripts != nil {
		if err := m.validateScripts(); err != nil {
			return err
		}
	}
	return nil
}

func (m Metadata) validateControl() error {
	if _, ok := m.Control["package"]; !ok {
		return errors.New("control:name is empty or missing")
	}
	if _, ok := m.Control["version"]; !ok {
		return errors.New("control:version is empty or missing")
	}
	if _, ok := m.Control["architecture"]; !ok {
		return errors.New("control:arch is empty or missing")
	}
	return nil
}

func (m Metadata) validateScripts() error {
	for name, content := range m.Scripts {
		if name == "" {
			return errors.New("script:name is empty or missing")
		}
		if content == "" {
			return errors.New("script:content is empty or missing")
		}
	}
	return nil
}

func (m Metadata) validateData() error {
	for src, dst := range m.Data {
		if src == "" {
			return errors.New("data:src is empty or missing")
		}
		if dst == "" {
			return errors.New("data:dst is empty or missing")
		}
	}
	return nil
}

func (m Metadata) Generate(dirname string) error {
	var err error
	defer func() {
		if err == nil {
			return
		}
		os.RemoveAll(dirname)
	}()
	err = m.generateControl(dirname)
	if err != nil {
		return err
	}
	err = m.generateData(dirname)
	if err != nil {
		return err
	}
	err = m.generateScripts(dirname)
	if err != nil {
		return err
	}
	return nil
}

func (m Metadata) generateControl(dirname string) error {
	debianDirPath := filepath.Join(dirname, "DEBIAN")
	if err := os.MkdirAll(debianDirPath, os.ModePerm); err != nil {
		return err
	}
	control, err := os.Create(filepath.Join(debianDirPath, "control"))
	if err != nil {
		return err
	}
	defer control.Close()
	for k, v := range m.Control {
		if _, err := control.WriteString(k + ": " + v + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func (m Metadata) generateData(dirname string) error {
	for src, dst := range m.Data {
		srcPath := filepath.Clean(src)
		fi, err := os.Stat(srcPath)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dirname, dst)
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

func (m Metadata) generateScripts(dirname string) error {
	debianDirPath := filepath.Join(dirname, "DEBIAN")
	for name, content := range m.Scripts {
		path := filepath.Join(debianDirPath, name)
		f, err := os.Create(path)
		if err != nil {
			return err
		}
		_, err = f.WriteString(content)
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
