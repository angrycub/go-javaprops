package javaprops

import (
	"bytes"
	"encoding/xml"
	"os"
	"os/exec"
	"path"

	"github.com/angrycub/go-javaprops/pkg/javaprops/internal/javasrc"
)

type runner struct {
	cfg jfConfig // a config so we can get the path under test
	td  string   // td is the temp dir the runner writes the java code to.
	pj  string   // full path to props.java
}

func getPropsFromRuntime(cfg jfConfig) (Props, error) {
	var err error
	r := &runner{
		cfg: cfg,
	}

	if r.td, err = os.MkdirTemp("", "nomad-javaprops"); err != nil {
		return nil, err
	}
	defer os.RemoveAll(r.td)

	tp := path.Join(r.td, "props.java")
	err = os.WriteFile(tp, javasrc.PropsJavaSrc, 0600)
	if err != nil {
		return nil, err
	}
	r.pj = tp

	return r.runPropsJava()
}

func (r runner) runPropsJava() (Props, error) {
	res := r.runJava([]string{r.pj})
	if res.err != nil {
		return nil, res.err
	}
	b := res.stdout.Bytes()
	pf := JVMPropertyFile{}
	if err := xml.Unmarshal(b, &pf); err != nil {
		return nil, err
	}
	return pf.Entries.AsProps(), nil
}

func (r runner) runJava(args []string) output {
	var out output

	jCmd := r.cfg.javaCmd
	if jCmd == "" {
		jCmd = "java"
	}

	// Fixup the command slice based on the runner config and the args passed to the func
	var javaCmd = append([]string{path.Join(r.cfg.jvmPath, jCmd)}, args...)

	cmd := exec.Command(javaCmd[0], javaCmd[1:]...)
	cmd.Stdout = &out.stdout
	cmd.Stderr = &out.stderr
	out.err = cmd.Run()
	return out
}

type output struct {
	stdout bytes.Buffer
	stderr bytes.Buffer
	err    error
}
