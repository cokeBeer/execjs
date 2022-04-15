package execjs

import (
	"reflect"
	"testing"
)

func TestExec_(t *testing.T) {
	need := 1.0
	output, _ := Exec_(`return 1`)
	res := output.(float64)
	if res != need {
		t.Fatalf("require %v, but get %v", need, res)
	}
}

func TestEval(t *testing.T) {
	need := "12"
	output, _ := Eval(`1+"2"`)
	res := output.(string)
	if res != need {
		t.Fatalf("require %s, but get %s", need, output)
	}
}

func TestEval2(t *testing.T) {
	need := 3.0
	output, _ := Eval(`1+2`)
	res := output.(float64)
	if res != need {
		t.Fatalf("require %v, but get %v", need, res)
	}
}

func TestEval3(t *testing.T) {
	need := []float64{1.0, 2.0, 4.0}
	output, _ := Eval(`[1,2,4]`)
	res := make([]float64, 0)
	for _, v := range output.([]interface{}) {
		res = append(res, v.(float64))
	}
	if !reflect.DeepEqual(res, need) {
		t.Fatalf("require %v, but get %v", need, res)
	}
}

func TestCall(t *testing.T) {
	need := 12.0
	c, _ := Compile(`function add(x, y) {
		return x.value + y;
	   }`)
	output, err := c.Call("add", map[string]float64{"value": 10}, 2)
	res := output.(float64)
	if res != need {
		t.Fatalf("require %v, but get %v, err:%s", need, res, err)
	}
}
