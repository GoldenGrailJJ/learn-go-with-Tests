package strings

import (
	"strings"
	"testing"
)

// 重复打印字符串
func TestStrRepeat(t *testing.T) {
	repeated := strings.Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

// 判断字符串是否包含某个子串
func TestStrCon(t *testing.T) {
	repeated := strings.Contains("a boy", "boy")
	expected := true

	if repeated != expected {
		t.Errorf("expected '%t' but got '%t'", expected, repeated)
	}
}

// 判断字符串是否以某个字符串为前缀
func TestHasPrefix(t *testing.T) {
	repeated := strings.HasPrefix("you are a boy", "you")
	expected := true

	if repeated != expected {
		t.Errorf("expected '%t' but got '%t'", expected, repeated)
	}
}

func TestSplit(t *testing.T) {
	repeated := strings.Split("a,b,c", ",")
	expected := [3]string{"a", "b", "c"}

	// 先判断长度是否一致
	if len(repeated) != len(expected) {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}

	// 长度一致在循环判断每个元素是否一致
	for i := 0; i < len(repeated); i++ {
		if repeated[i] != expected[i] {
			t.Errorf("expected '%q' but got '%q'", expected, repeated)
			break
		}
	}

}
