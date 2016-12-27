package server_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/go-microservices/resizer/option"
	"github.com/go-microservices/resizer/server"
)

var (
	appServer      *httptest.Server
	fixturesServer *httptest.Server
)

func TestMain(m *testing.M) {
	fixturesServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		p := path.Join(dir, "..", "fixtures", r.URL.Path[1:])
		http.ServeFile(w, r, p)
	}))
	u, err := url.Parse(fixturesServer.URL)
	if err != nil {
		panic(err)
	}

	h, err := server.NewHandler(option.Options{
		DBUser:     "root",
		DBProtocol: "tcp",
		DBAddress:  "mysql:3306",
		DBName:     "resizer",
		JSON:       "/secret/gcloud.json",
		Hosts:      []string{u.Host},
	})
	if err != nil {
		panic(err)
	}
	appServer = httptest.NewServer(http.HandlerFunc(h.ServeHTTP))

	c := m.Run()

	appServer.Close()
	fixturesServer.Close()

	os.Exit(c)
}

func TestNew(t *testing.T) {
	func() {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				t.Error("request at the 1st time shouldn't be redirected")
				return nil
			},
		}
		resp, err := client.Get(fmt.Sprintf("%s?width=15&url=%s/f-png24.png", appServer.URL, fixturesServer.URL))
		if err != nil {
			t.Fatalf("fail to get resized image at the 1st time: error=%v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("1st time should response as ok: expected %d, but actual %d", http.StatusOK, resp.StatusCode)
		}
	}()

	time.Sleep(time.Second * 5)

	func() {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if strings.Contains(req.URL.Host, "googleapis") != true {
					t.Error("request at the 2nd time shouldn't be redirected")
				}
				return nil
			},
		}
		resp, err := client.Get(fmt.Sprintf("%s?width=15&url=%s/f-png24.png", appServer.URL, fixturesServer.URL))
		if err != nil {
			t.Fatalf("fail to get resized image at the 2st time: error=%v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("response at the 2nd time should be ok: expected %d, but actual %d", http.StatusOK, resp.StatusCode)
		}
	}()

	func() {
		client := &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				if strings.Contains(req.URL.Host, "googleapis") != true {
					t.Error("request at the 3rd time shouldn't be redirected")
				}
				return nil
			},
		}
		resp, err := client.Get(fmt.Sprintf("%s?height=21&url=%s/f-png24.png", appServer.URL, fixturesServer.URL))
		if err != nil {
			t.Errorf("fail to get resized image at the 3rd time: error=%v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("response at the 3rd time should be ok: expected %d, but actual %d", http.StatusOK, resp.StatusCode)
		}
	}()
}

func TestFail(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("%s", appServer.URL))
	if err != nil {
		t.Fatalf("fail to get resized image at the 1st time: error=%v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("should be fail")
	}
	dec := json.NewDecoder(resp.Body)
	var v map[string]string
	dec.Decode(&v)
	if v["message"] == "" {
		t.Errorf("should have a message")
	}
}
