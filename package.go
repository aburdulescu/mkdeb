package main

import "errors"

type Package struct {
	Control   map[string]string `json:"control"`
	Scripts   map[string]string `json:"scripts,omitempty"`
	Artifacts map[string]string `json:"artifacts"`
}

func (p Package) Validate() error {
	if err := p.validateControl(); err != nil {
		return errors.New("package:" + err.Error())
	}
	if p.Artifacts == nil {
		return errors.New("package: artifacts is empty or missing")
	}
	if err := p.validateArtifacts(); err != nil {
		return errors.New("package:" + err.Error())
	}
	if p.Scripts != nil {
		if err := p.validateScripts(); err != nil {
			return errors.New("package:" + err.Error())
		}
	}
	return nil
}

func (p Package) validateControl() error {
	if _, ok := p.Control["package"]; !ok {
		return errors.New("control:name is empty or missing")
	}
	if _, ok := p.Control["version"]; !ok {
		return errors.New("control:version is empty or missing")
	}
	if _, ok := p.Control["architecture"]; !ok {
		return errors.New("control:arch is empty or missing")
	}
	return nil
}

func (p Package) validateScripts() error {
	for name, content := range p.Scripts {
		if name == "" {
			return errors.New("script:name is empty or missing")
		}
		if content == "" {
			return errors.New("script:content is empty or missing")
		}
	}
	return nil
}

func (p Package) validateArtifacts() error {
	for src, dst := range p.Artifacts {
		if src == "" {
			return errors.New("artifact:src is empty or missing")
		}
		if dst == "" {
			return errors.New("artifact:dst is empty or missing")
		}
	}
	return nil
}
