package askpass

import (
	"errors"
	"os"
	"testing"
)

func testPassword(t *testing.T) string {
	t.Helper()
	return "pass2"
}

func setup(t *testing.T) *Pass {
	t.Helper()
	filename := "test.pass"

	pass := &Pass{Filename: filename}
	_ = os.Remove(filename)

	return pass
}

func teardown(t *testing.T, filename string) {
	t.Helper()
	_ = os.Remove(filename)
}

func Test_GetPass_and_SavePass(t *testing.T) {
	pass := setup(t)
	defer teardown(t, pass.Filename)

	err := pass.Save(testPassword(t))
	if err != nil {
		t.Fatalf("Save() failed, error = %v", err)
	}

	got, err := pass.Get()
	if err != nil {
		t.Fatalf("Get() failed, error = %v", err)
	}

	if got != testPassword(t) {
		t.Errorf("Get() = %q, want %q", got, testPassword(t))
	}
}

func TestSaveAndRetrievePassword(t *testing.T) {
	pass := setup(t)
	defer teardown(t, pass.Filename)

	err := pass.Save(testPassword(t))
	if err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	savedPassword, err := pass.Get()
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}

	if savedPassword != testPassword(t) {
		t.Errorf("Get() = %v, want %v", savedPassword, testPassword(t))
	}
}

func TestFileNotFound(t *testing.T) {
	pass := setup(t)
	defer teardown(t, pass.Filename)

	_, err := pass.Get()
	if err == nil || !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Expected file not exist error, got %v", err)
	}
}

func TestEmptyPassword(t *testing.T) {
	pass := setup(t)
	defer teardown(t, pass.Filename)

	err := pass.Save("")
	if err == nil {
		t.Fatalf("Expected error for empty password, got nil")
	}
}

func TestFilePermissions(t *testing.T) {
	pass := setup(t)
	defer teardown(t, pass.Filename)

	err := pass.Save(testPassword(t))
	if err != nil {
		t.Fatalf("Save() error = %v", err)
	}

	info, err := os.Stat(pass.Filename)
	if err != nil {
		t.Fatalf("Stat() error = %v", err)
	}

	if info.Mode().Perm() != 0600 {
		t.Errorf("File permissions = %v, want %v", info.Mode().Perm(), 0600)
	}
}
