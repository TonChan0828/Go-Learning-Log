package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_parser(t *testing.T) {
	tests := []struct {
		name    string
		data    string
		want    Input
		wantErr bool
	}{
		{
			name: "valid input",
			data: "CALC_1\n+\n3\n2",
			want: Input{
				Id:   "CALC_1",
				Op:   "+",
				Val1: 3,
				Val2: 2,
			},
		},
		{
			name:    "bad number",
			data:    "CALC_BAD\n+\nthree\n2",
			wantErr: true,
		},
		{
			name:    "bad second number",
			data:    "CALC_BAD2\n+\n1\ntwo",
			wantErr: true,
		},
		{
			name:    "missing fields",
			data:    "ONLY_ID\n+\n1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser([]byte(tt.data))
			if (err != nil) != tt.wantErr {
				t.Fatalf("parser() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				return
			}
			if got != tt.want {
				t.Fatalf("parser() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func TestDataProcessor(t *testing.T) {
	in := make(chan []byte, 7)
	out := make(chan Result, 7)

	// feed a mix of valid and invalid payloads
	inputs := []string{
		"ID_PLUS\n+\n1\n2",
		"ID_MINUS\n-\n5\n3",
		"ID_MULT\n*\n4\n10",
		"ID_DIV\n/\n8\n2",
		"ID_DIVZERO\n/\n5\n0",
		"ID_BADOP\n?\n1\n1",
		"ID_BADNUM\n+\none\n1",
	}

	for _, payload := range inputs {
		in <- []byte(payload)
	}
	close(in)

	go DataProcessor(in, out)

	var got []Result
	for r := range out {
		got = append(got, r)
	}

	want := []Result{
		{Id: "ID_PLUS", Value: 3},
		{Id: "ID_MINUS", Value: 2},
		{Id: "ID_MULT", Value: 40},
		{Id: "ID_DIV", Value: 4},
	}

	if len(got) != len(want) {
		t.Fatalf("unexpected result count: got %d want %d", len(got), len(want))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("result[%d] = %#v, want %#v", i, got[i], want[i])
		}
	}
}

func TestWriteData(t *testing.T) {
	buf := bytes.Buffer{}
	out := make(chan Result, 2)
	out <- Result{Id: "ID1", Value: 10}
	out <- Result{Id: "ID2", Value: 20}
	close(out)

	WriteData(out, &buf)

	want := "ID1:10\nID2:20\n"
	if buf.String() != want {
		t.Fatalf("WriteData output = %q, want %q", buf.String(), want)
	}
}

func TestNewController(t *testing.T) {
	queue := make(chan []byte, 1)
	handler := NewController(queue)

	req1 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("CALC_1\n+\n1\n1"))
	w1 := httptest.NewRecorder()
	handler.ServeHTTP(w1, req1)

	if w1.Code != http.StatusAccepted {
		t.Fatalf("first request status = %d, want %d", w1.Code, http.StatusAccepted)
	}
	if body := strings.TrimSpace(w1.Body.String()); body != "OK: 1" {
		t.Fatalf("first response body = %q, want %q", body, "OK: 1")
	}
	if len(queue) != 1 {
		t.Fatalf("queue length after first request = %d, want %d", len(queue), 1)
	}

	req2 := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("CALC_2\n+\n2\n2"))
	w2 := httptest.NewRecorder()
	handler.ServeHTTP(w2, req2)

	if w2.Code != http.StatusServiceUnavailable {
		t.Fatalf("second request status = %d, want %d", w2.Code, http.StatusServiceUnavailable)
	}
	if body := strings.TrimSpace(w2.Body.String()); body != "Too Busy: 1" {
		t.Fatalf("second response body = %q, want %q", body, "Too Busy: 1")
	}
}

type errReadCloser struct{}

func (e errReadCloser) Read(_ []byte) (int, error) { return 0, errors.New("read error") }
func (e errReadCloser) Close() error               { return nil }

func TestNewControllerBadInput(t *testing.T) {
	queue := make(chan []byte, 1)
	handler := NewController(queue)

	req := httptest.NewRequest(http.MethodPost, "/", errReadCloser{})
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("bad input status = %d, want %d", w.Code, http.StatusBadRequest)
	}
	if body := strings.TrimSpace(w.Body.String()); body != "Bad Input" {
		t.Fatalf("bad input response body = %q, want %q", body, "Bad Input")
	}
}

func Test_runServer(t *testing.T) {
	origListen := listenAndServe
	origCreate := createFile
	defer func() {
		listenAndServe = origListen
		createFile = origCreate
	}()

	listenCalled := false
	listenAndServe = func(addr string, handler http.Handler) error {
		listenCalled = true
		if addr != ":8080" {
			t.Fatalf("listen address = %q, want %q", addr, ":8080")
		}
		// ensure handler responds
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("ID\n+\n1\n1"))
		rec := httptest.NewRecorder()
		handler.ServeHTTP(rec, req)
		if rec.Code != http.StatusAccepted {
			t.Fatalf("handler status = %d, want %d", rec.Code, http.StatusAccepted)
		}
		return nil
	}

	var tmpPath string
	createFile = func(name string) (*os.File, error) {
		f, err := os.CreateTemp(".", "results-test-")
		if err != nil {
			return nil, err
		}
		tmpPath = f.Name()
		return f, nil
	}

	if err := runServer(); err != nil {
		t.Fatalf("runServer returned error: %v", err)
	}

	if !listenCalled {
		t.Fatalf("listenAndServe was not called")
	}
	if tmpPath == "" {
		t.Fatalf("createFile stub was not used")
	}
	os.Remove(tmpPath)
}

func Test_runServerCreateFileError(t *testing.T) {
	origCreate := createFile
	defer func() { createFile = origCreate }()

	expected := errors.New("create failed")
	createFile = func(name string) (*os.File, error) {
		return nil, expected
	}

	if err := runServer(); !errors.Is(err, expected) {
		t.Fatalf("runServer error = %v, want %v", err, expected)
	}
}

func Test_mainErrorPath(t *testing.T) {
	origRun := runApp
	defer func() { runApp = origRun }()

	runApp = func() error { return errors.New("boom") }

	origStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = origStdout
	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "boom") {
		t.Fatalf("expected error output to contain %q, got %q", "boom", string(out))
	}
}

func FuzzParser(f *testing.F) {
	// Seed with representative payloads
	f.Add("ID\n+\n1\n2")
	f.Add("ONLY_ID\n+\n1")
	f.Add("ID\n-\n10\n5")
	f.Add("ID\n/\n10\n0")

	f.Fuzz(func(t *testing.T, payload string) {
		input, err := parser([]byte(payload))
		if err != nil {
			return
		}
		// If parsing succeeds, re-encode and parse again; must match.
		roundTrip := fmt.Sprintf("%s\n%s\n%d\n%d", input.Id, input.Op, input.Val1, input.Val2)
		input2, err := parser([]byte(roundTrip))
		if err != nil {
			t.Fatalf("round-trip parse failed: %v", err)
		}
		if input != input2 {
			t.Fatalf("round-trip mismatch: %#v != %#v", input, input2)
		}
	})
}
