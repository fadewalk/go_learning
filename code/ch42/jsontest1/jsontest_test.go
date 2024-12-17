package jsontest

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
type JobInfo struct {
	Skills []string `json:"skills"`
}
type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}

var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}`

// 辅助函数：打印结构体的键和值
func printStructFields(obj interface{}) {
	fmt.Println("Object:", obj)
	val := reflect.ValueOf(obj)
	fmt.Println("Value:", val)
	typ := val.Type()

	fmt.Println("Type:", typ)

	fmt.Println("NumField:", val.NumField())
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := typ.Field(i)
		fmt.Printf("%s: %v\n", fieldType.Name, fieldVal.Interface())
	}
}

func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("JobInfo:")
	printStructFields(e.JobInfo)
	fmt.Println("BasicInfo:")
	printStructFields(e.BasicInfo)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}

	fmt.Printf("e %v", *e)
	fmt.Printf("e %+v", *e)
	fmt.Printf("e %#v", *e)
}
