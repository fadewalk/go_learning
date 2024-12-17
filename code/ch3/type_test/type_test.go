package type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64

	//b = a  // 不可以从小类型范围到大的转换

	b = int64(a)
	var c MyInt
	//c = b // 即使是别名，也不可以转换
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1  // 不能运算指针
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") //初始化零值是空字符串 ""，而不是nil
	t.Log(len(s))

}
