package execjs

import (
	"fmt"
	"strings"
)

const (
	Node           string = "Node"
	JavaScriptCore string = "JavaScriptCore"
	SpiderMonkey   string = "SpiderMonkey"
	JScript        string = "JScript"
	PhantomJS      string = "PhantomJS"
	SlimerJS       string = "SlimerJS"
	Nashorn        string = "Nashorn"
)

type Tuple struct {
	Name    string
	Runtime RuntimeInterface
}

var runtimes []Tuple

func init() {
	Register(Node, node())
}

func Register(name string, runtime RuntimeInterface) {
	runtimes = append(runtimes, Tuple{Name: name, Runtime: runtime})
}

func GetRuntime(name string) (RuntimeInterface, error) {
	if name == "" {
		runtime, err := find_available_runtime()
		if err != nil {
			return nil, err
		} else {
			return runtime, nil
		}
	}
	runtime, err := find_runtime_by_name(name)
	if err != nil {
		return nil, err
	} else {
		return runtime, nil
	}
}

func Runtimes() []Tuple {
	return runtimes
}

func find_available_runtime() (RuntimeInterface, error) {
	var runtime RuntimeInterface
	runtime = nil
	for _, tuple := range runtimes {
		if tuple.Runtime.Is_available() {
			runtime = tuple.Runtime
		}
	}
	if runtime != nil {
		return runtime, nil
	} else {
		return nil, RuntimeUnavailableError{Message: "Could not find an available JavaScript runtime."}
	}
}

func find_runtime_by_name(name string) (RuntimeInterface, error) {
	var runtime RuntimeInterface
	runtime = nil
	for _, tuple := range runtimes {
		if strings.ToLower(tuple.Name) == strings.ToLower(name) {
			runtime = tuple.Runtime
			break
		}
	}
	if runtime != nil {
		if runtime.Is_available() {
			return runtime, nil
		} else {
			return nil, RuntimeUnavailableError{Message: fmt.Sprintf("%s runtime is not available on this system", name)}
		}
	} else {
		return nil, RuntimeUnavailableError{Message: fmt.Sprintf("%s runtime is not defined", name)}
	}
}

func Eval(source string) (interface{}, error) {
	runtime, err := GetRuntime("")
	if err != nil {
		return "", err
	}
	return runtime.Eval(source)
}

func Exec_(source string) (interface{}, error) {
	runtime, err := GetRuntime("")
	if err != nil {
		return "", err
	}
	return runtime.Exec_(source)
}

func Compile(source string) (RuntimeContextInterface, error) {
	runtime, err := GetRuntime("")
	if err != nil {
		return nil, err
	}
	return runtime.Compile(source), nil
}
