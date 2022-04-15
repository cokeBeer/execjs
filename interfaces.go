package execjs

type RuntimeInterface interface {
	Exec_(string) (interface{}, error)
	Eval(string) (interface{}, error)
	Compile(string) RuntimeContextInterface
	Is_available() bool
	compile(string) RuntimeContextInterface
}

type RuntimeContextInterface interface {
	Exec_(string) (interface{}, error)
	Eval(string) (interface{}, error)
	Call(string, ...interface{}) (interface{}, error)
	Is_available() bool
	exec_(string) (interface{}, error)
	eval(string) (interface{}, error)
	call(string, ...interface{}) (interface{}, error)
}
