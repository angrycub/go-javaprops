package javaprops

import "fmt"

type jfConfig struct {
	jvmPath string // defaults to nothing when unset so the system path is used
	javaCmd string // defaults to "java" when unset
}

type jfOption func(c jfConfig) jfConfig

func WithJavaPath(p string) jfOption {
	return func(c jfConfig) jfConfig {
		c.jvmPath = p
		return c
	}
}

func WithJavaCmd(cmd string) jfOption {
	return func(c jfConfig) jfConfig {
		c.javaCmd = cmd
		return c
	}
}

func New(opts ...jfOption) jfConfig {
	c := jfConfig{}
	for _, opt := range opts {
		c = opt(c)
	}
	return c
}

func (j jfConfig) ErrNoJava() error {
	pTxt := "at " + j.jvmPath
	if j.jvmPath == "" {
		pTxt = "on system path"
	}
	return fmt.Errorf("no %s executable found %s", j.javaCmd, pTxt)
}

func (j jfConfig) GetProps() (Props, error) {
	return getPropsFromRuntime(j)
}
