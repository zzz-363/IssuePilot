package config

import "testing"

// 测试默认端口
func TestGetPort_Default(t *testing.T) {
	// 清除 PORT 环境变量
	t.Setenv("PORT", "")

	got := GetPort()
	want := "8080"

	if got != want {
		t.Errorf("GetPort() = %q, want %q", got, want)
	}
}

// 测试自定义端口
func TestGetPort_Custom(t *testing.T) {
	t.Setenv("PORT", "9000")

	got := GetPort()
	want := "9000"

	if got != want {
		t.Errorf("GetPort() = %q, want %q", got, want)
	}
}
