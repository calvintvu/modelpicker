package inference

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestRuntimeLibraryPath(t *testing.T) {
	t.Setenv("ONNX_RUNTIME_PATH", "")
	for _, tc := range []struct {
		goos, goarch, want string
	}{
		{"linux", "amd64", "lib/linux-amd64/lib/libonnxruntime.so"},
		{"linux", "arm64", "lib/linux-arm64/lib/libonnxruntime.so"},
		{"darwin", "arm64", "lib/darwin-arm64/lib/libonnxruntime.dylib"},
		{"windows", "amd64", "lib/windows-amd64/lib/onnxruntime.dll"},
	} {
		t.Run(tc.goos+"-"+tc.goarch, func(t *testing.T) {
			if got := runtimeLibraryPath(tc.goos, tc.goarch); got != filepath.FromSlash(tc.want) {
				t.Fatalf("got %q, want %q", got, filepath.FromSlash(tc.want))
			}
		})
	}
}

func TestLoadInferenceSessionDefaultRuntimePath(t *testing.T) {
	runtimePath := os.Getenv("TEST_ONNX_RUNTIME_PATH")
	if runtimePath == "" {
		t.Skip("set TEST_ONNX_RUNTIME_PATH to test native runtime loading")
	}
	runtimePath, err := filepath.Abs(runtimePath)
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	name := "libonnxruntime.so"
	if runtime.GOOS == "darwin" {
		name = "libonnxruntime.dylib"
	}
	path := filepath.Join(dir, "lib", runtime.GOOS+"-"+runtime.GOARCH, "lib", name)
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(runtimePath, path); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)
	t.Setenv("ONNX_RUNTIME_PATH", "")
	t.Setenv("MODEL_DIR", "")
	session, err := LoadInferenceSession()
	if session != nil {
		session.CloseInferenceSession()
	}
	if err == nil || !strings.Contains(err.Error(), "loading tokenizer") ||
		!strings.Contains(err.Error(), filepath.Join("inference", "model", "tokenizer", "tokenizer.json")) {
		t.Fatalf("expected runtime initialization to succeed before missing default tokenizer, got %v", err)
	}
}

func TestLoadInferenceSessionRuntimePathOverride(t *testing.T) {
	path := filepath.Join(t.TempDir(), "custom runtime", "missing-library")
	t.Setenv("ONNX_RUNTIME_PATH", path)
	session, err := LoadInferenceSession()
	if session != nil {
		session.CloseInferenceSession()
	}
	if err == nil || !strings.Contains(err.Error(), path) {
		t.Fatalf("expected initialization error naming %q, got %v", path, err)
	}
}

func TestLoadInferenceSessionModelDir(t *testing.T) {
	runtimePath := os.Getenv("TEST_ONNX_RUNTIME_PATH")
	if runtimePath == "" {
		t.Skip("set TEST_ONNX_RUNTIME_PATH to test model paths with the native runtime")
	}
	for _, withTokenizer := range []bool{false, true} {
		name := "missing tokenizer"
		if withTokenizer {
			name = "missing model"
		}
		t.Run(name, func(t *testing.T) {
			dir := t.TempDir()
			t.Setenv("ONNX_RUNTIME_PATH", runtimePath)
			t.Setenv("MODEL_DIR", dir)
			want := filepath.Join(dir, "tokenizer", "tokenizer.json")
			if withTokenizer {
				if err := os.Mkdir(filepath.Dir(want), 0755); err != nil {
					t.Fatal(err)
				}
				data := `{"version":"1.0","model":{"type":"WordLevel","vocab":{"[UNK]":0,"hello":1},"unk_token":"[UNK]"}}`
				if err := os.WriteFile(want, []byte(data), 0644); err != nil {
					t.Fatal(err)
				}
				want = filepath.Join(dir, "laya.onnx")
			}
			session, err := LoadInferenceSession()
			if session != nil {
				session.CloseInferenceSession()
			}
			if err == nil || !strings.Contains(err.Error(), want) {
				t.Fatalf("expected loading error naming %q, got %v", want, err)
			}
		})
	}
}
