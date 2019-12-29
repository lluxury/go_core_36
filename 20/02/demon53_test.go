package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	var name string
	greeting, err := hello(name)
	if err == nil {		// 无错,无值检测
		t.Errorf("The error is nil, but it should not be. (name=%q)",
			name)
	}

	if greeting != "" {   // 空
		t.Errorf("Nonempty greeting, but it should not be. (name=%q)",
			name)
	}

	name = "Robert"
	greeting, err = hello(name)
	if err != nil {    // 有错
		t.Errorf("The error is not nil, but it should be. (name=%q)",
			name)
	}

	if greeting == "" {   // 无值
		t.Errorf("Empty greeting, but it should not be. (name=%q)",
			name)
	}
	expected := fmt.Sprintf("Hello, %s!", name)
	if greeting != expected {    // 非期望值
		t.Errorf("The actual greeting %q is not the expected. (name=%q)",
			greeting, name)
	}
	t.Logf("The expected greeting is %q.\n", expected)
}

func testIntroduce(t *testing.T) {
	intro := introduce()
	expected := "Welcome to my Golang column."  // 匹配值
	if intro != expected {
		t.Errorf("The actual introduce %q is not the expected.",
			intro)
	}
	t.Logf("The expected introduce is %q.\n", expected)
}

func TestFail(t *testing.T)  {
	//t.Fail()
	t.FailNow()
	t.Log("Failed.")
}