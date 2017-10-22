package statink2

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNewClient(t *testing.T) {
	dummyKey := "APIKEY"
	_, err := NewClient(dummyKey, nil)
	if err != nil {
		t.Errorf("NewClient() should not be error with %v", err)
	}
}

func TestPostBattle(t *testing.T) {
	dummyKey := "APIKEY"

	muxAPI := http.NewServeMux()
	muxAPI.HandleFunc("/api/v2/battle", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("", "")
		strs := strings.Split(r.Header.Get("Authorization"), " ")
		if strs[1] != dummyKey {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"error" : "invalid api key"}`))
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	testServer := httptest.NewServer(muxAPI)
	defer testServer.Close()

	c, err := newClient(testServer.URL, dummyKey, nil)
	if err != nil {
		t.Errorf("newClient() should not be error with err : %v", err)
	}
	b := Battle{}
	err = c.PostBattle(nil, b)
	if err != nil {
		t.Errorf("PostBattle() should not be error with err : %v", err)
	}
}
