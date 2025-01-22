package execute

import (
	"os/exec"
	"path/filepath"
	"strings"
)

type BinArg string
type BinPath string

const (
	HTMLtoIMG BinPath = "./wkhtml/wkhtmltoimage"
	HTMLtoPDF BinPath = "./wkhtml/wkhtmltopdf"
)

func HTMLtoIMGPath() BinPath {
	return "./" + BinPath(filepath.Clean(string(HTMLtoIMG)))
}

func HTMLtoPDFPath() BinPath {
	return "./" + BinPath(filepath.Clean(string(HTMLtoIMG)))
}

type Executer struct {
	bin  BinPath
	args []BinArg
}

func NewExecuter(bin BinPath, args []BinArg) *Executer {
	return &Executer{
		bin:  bin,
		args: args,
	}
}
func (e *Executer) ProcessConverter(content []byte) (out []byte, err error) {
	binCmd := exec.Command(string(e.bin), e.ArgsString(), "-", "-")
	binCmd.Stdin = strings.NewReader(string(content))

	out, err = binCmd.CombinedOutput()
	return
}

func (e *Executer) ArgsString() string {
	out := make([]string, len(e.args))

	for i, arg := range e.args {
		out[i] = "-" + string(arg)
	}

	return strings.Join(out, " ")
}
