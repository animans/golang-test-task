package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang-test-task/internal/app/storage"
)

type mockRepo struct {
	addFn  func(ctx context.Context, v int64) error
	listFn func(ctx context.Context) ([]int64, error)
}

func (m mockRepo) Add(ctx context.Context, v int64) error          { return m.addFn(ctx, v) }
func (m mockRepo) ListSorted(ctx context.Context) ([]int64, error) { return m.listFn(ctx) }

var _ storage.Repository = (*mockRepo)(nil)

func TestAddNumber_ReturnsSortedList(t *testing.T) {
	repo := mockRepo{
		addFn: func(ctx context.Context, v int64) error { return nil },
		listFn: func(ctx context.Context) ([]int64, error) {
			return []int64{1, 2, 3}, nil
		},
	}

	h := NewHandler(repo)
	srv := httptest.NewServer(Router(h))
	t.Cleanup(srv.Close)

	body, _ := json.Marshal(AddNumberRequest{Value: 3})
	resp, err := http.Post(srv.URL+"/numbers", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var out NumbersResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		t.Fatal(err)
	}

	want := []int64{1, 2, 3}
	if len(out.Numbers) != len(want) {
		t.Fatalf("unexpected len: %v", out.Numbers)
	}
	for i := range want {
		if out.Numbers[i] != want[i] {
			t.Fatalf("expected %v, got %v", want, out.Numbers)
		}
	}
}
