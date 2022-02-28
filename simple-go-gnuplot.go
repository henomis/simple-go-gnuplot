package simplegognuplot

import (
	"fmt"
	"os"
	"strings"

	"github.com/mibk/shellexec"
)

type SimpleGnuplot struct {
	gnuplotExecutablePath string
	environmentVariables  map[string]interface{}
	plotFilePath          string
}

func New(gnuplotExecutablePath string) *SimpleGnuplot {

	return &SimpleGnuplot{
		gnuplotExecutablePath: gnuplotExecutablePath,
		environmentVariables:  make(map[string]interface{}),
	}

}

func (sg *SimpleGnuplot) AddEnv(key string, value interface{}) {
	sg.environmentVariables[key] = value
}

func (sg *SimpleGnuplot) DeleteEnv() {
	sg.environmentVariables = make(map[string]interface{})
}

func (sg *SimpleGnuplot) SetPlotFilePath(plotFilePath string) {
	sg.plotFilePath = plotFilePath
}

func (sg *SimpleGnuplot) Exec() error {

	err := sg.validate()
	if err != nil {
		return err
	}

	command, err := shellexec.Command(sg.buildGnuplotCommand())

	if err != nil {
		return err
	}

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		return err
	}

	return command.Wait()

}

// SUPPORT METHODS

func (sg *SimpleGnuplot) validate() error {

	fileInfo, err := os.Stat(sg.gnuplotExecutablePath)
	if err != nil {
		return err
	}

	if (fileInfo.Mode() & 0111) == 0 {
		return fmt.Errorf("permission denied")
	}

	_, err = os.Stat(sg.plotFilePath)
	if err != nil {
		return err
	}

	return nil

}

func (sg *SimpleGnuplot) buildGnuplotCommand() string {

	gnuplotCommand := sg.gnuplotExecutablePath

	if len(sg.environmentVariables) > 0 {
		gnuplotCommand = fmt.Sprintf("%s -e \"%s\"", gnuplotCommand, sg.buildGnuplotEnvironmentList())
	}

	gnuplotCommand = fmt.Sprintf("%s %s", gnuplotCommand, sg.plotFilePath)

	return gnuplotCommand

}

func (sg *SimpleGnuplot) buildGnuplotEnvironmentList() string {

	gnuplotEnvironmentList := ""

	for key, value := range sg.environmentVariables {
		gnuplotEnvironmentList += fmt.Sprintf("%s='%v'; ", key, value)
	}

	return strings.TrimSuffix(gnuplotEnvironmentList, "; ")

}
