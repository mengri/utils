package autowire

import (
	"fmt"
	"testing"
)

type ITest interface {
	Test()
	Name() string
}
type TestIml1 struct {
	name   string
	tester ITest `autowired:""`
}

func (t *TestIml1) OnCreate() {
	fmt.Printf("%s:on create\n", t.name)
}

func (t *TestIml1) Initialized() {
	fmt.Printf("%s:on initialized\n", t.name)

}

func (t *TestIml1) Test() {
	fmt.Printf("%s:test\n", t.name)
}

func (t *TestIml1) Name() string {
	return t.name
}

func (t *TestIml1) OnComplete() {
	fmt.Printf("%s:on complete\n", t.name)

}

func TestAuto(t *testing.T) {

	Auto(func() *TestIml1 {
		return &TestIml1{
			name: "Test1",
		}
	}, TypeName[TestIml1](), TypeName[ITest]())
	Auto(func() *int {
		i := new(int)
		*i = 99
		return i
	}, "system.port")
	var ti ITest
	var t2 *TestIml1
	var t3 = &TestIml1{
		name: "Test3",
	}
	Autowired(&ti)
	Autowired(&t2)
	Autowired(t3)
	Check()

	if ti != nil {
		ti.Test()
	}

	if t2 != nil {
		t2.Test()
		if t2.tester != nil {
			t2.tester.Test()
		}
	}

	if t3 != nil {
		t3.Test()
		if t3.tester != nil {
			t3.tester.Test()
		}
	}
}
