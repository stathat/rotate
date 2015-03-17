package rotate

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	root, err := ioutil.TempDir("", "multitest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(root)

	x, err := New(root, "mt")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := x.Write([]byte("hello\n")); err != nil {
		t.Fatal(err)
	}

	d, err := os.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	names, err := d.Readdirnames(1024)
	if err != nil {
		t.Fatal(err)
	}
	if len(names) != 1 {
		t.Errorf("number files in root: %d, expected 1", len(names))
	}
}

func TestRotate(t *testing.T) {
	root, err := ioutil.TempDir("", "multitest")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(root)

	x, err := New(root, "mt")
	if err != nil {
		t.Fatal(err)
	}
	x.SetMax(5)
	for i := 0; i < 20; i++ {
		if _, err := x.Write([]byte("hello\n")); err != nil {
			t.Fatal(err)
		}
	}

	d, err := os.Open(root)
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	names, err := d.Readdirnames(1024)
	if err != nil {
		t.Fatal(err)
	}
	if len(names) != 11 {
		t.Errorf("number files in root: %d, expected 11", len(names))
		for i, n := range names {
			t.Logf("%d: %q", i, n)
		}
	}
}
