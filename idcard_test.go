package idcard

import (
	"fmt"
	"testing"
)

func Benchmark_Upgrade15To18_Success(b *testing.B) {
	for index := 0; index < b.N; index++ {
		Upgrade15To18("370986890623212")
	}
}

func Test_Upgrade15To18_Success(t *testing.T) {
	newID, newIDErr := Upgrade15To18("370986890623212")

	if newIDErr != nil {
		t.Fatal(newIDErr)
		return
	}

	if newID != "370986198906232123" {
		t.Fatal("升位错误")
	}
}

func Test_Upgrade15To18_Not15(t *testing.T) {
	id := "370986890623212X"
	_, newIDErr := Upgrade15To18(id)

	if newIDErr != nil && newIDErr.Error() == fmt.Errorf("len(%s) is not 15", id).Error() {
		return
	}

	t.Fatal("非15位身份证的输入应该返货错误!")
}
