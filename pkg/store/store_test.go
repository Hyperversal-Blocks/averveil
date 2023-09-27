package store

import (
	"context"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	storer, _ := New(context.Background(), nil,
		"../../data/store", true, false)

	t.Run("ok", func(t *testing.T) {
		err := storer.Put("testKey", []byte("some test"))
		if err != nil {
			t.Fatal(err)
		}

		data, err := storer.Get("testKey")
		if err != nil {
			t.Fatal(err)
		}

		if strings.Compare(string(data), "some test") != 0 {
			if err != nil {
				t.Fatal(err)
			}
		}
	})
}
