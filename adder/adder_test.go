package integers

import "testing"
import "fmt"
/*
* 编写最小量通过测试代码，不是一次性把代码完成。
* go doc和example的编写
* 在函数已经清晰显示出含义的情况下，不使用命名返回值
* 在有example的测试代码中，使用go test -v加上对样例代码的测试
*/
func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5
	
	if expected != sum {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(2, 4)
	fmt.Println(sum)
	// Output: 6
}