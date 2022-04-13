package execjs

import (
	"testing"
)

func TestExec_(t *testing.T) {
	need := "1"
	output, _ := Exec_(`return 1`)
	if output != need {
		t.Fatalf("require %s, but get %s", need, output)
	}
}

func TestEval(t *testing.T) {
	need := `"12"`
	output, _ := Eval(`1+"2"`)
	if output != need {
		t.Fatalf("require %s, but get %s", need, output)
	}
}

func TestEval2(t *testing.T) {
	need := `3`
	output, _ := Eval(`1+2`)
	if output != need {
		t.Fatalf("require %s, but get %s", need, output)
	}
}

func TestEval3(t *testing.T) {
	need := `[1,2,4]`
	output, _ := Eval(`[1,2,4]`)
	if output != need {
		t.Fatalf("require %s, but get %s", need, output)
	}
}

func TestCall(t *testing.T) {
	need := `3`
	c, _ := Compile(`function add(x, y) {
		return x + y;
	   }`)
	output, _ := c.Call("add", `1`, `2`)
	if output != need {
		t.Fatalf("require %s, but get %s", need, output)
	}
}
