package execjs

type RuntimeInterface interface {
	Exec_(string) (string, error)
	Eval(string) (string, error)
	Compile(string) RuntimeContextInterface
	Is_available() bool
	compile(string) RuntimeContextInterface
}

type RuntimeContextInterface interface {
	Exec_(string) (string, error)
	Eval(string) (string, error)
	Call(string, ...string) (string, error)
	Is_available() bool
	exec_(string) (string, error)
	eval(string) (string, error)
	call(string, ...string) (string, error)
}
