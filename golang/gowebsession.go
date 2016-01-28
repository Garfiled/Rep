package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type TSession struct {
}

type TUMux struct {
	Sessions map[string]*TSession
	M        *http.ServeMux
	sync.Mutex
}

func (t *TUMux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("mysid")
	m := make(map[string]interface{})
	if err != nil {
		if req.URL.Path == "/login" {
			sid := GenSessionId()
			t.Lock()
			t.Sessions[sid] = &TSession{}
			t.Unlock()
			cookie := &http.Cookie{Name: "mysid", Value: sid,
				Expires: time.Now().Add(time.Hour), HttpOnly: true, MaxAge: 50000, Path: "/"}
			http.SetCookie(w, cookie)
			m["MsgCode"] = 1000
			m["Sid"] = sid
			bytes, _ := json.MarshalIndent(m, "", "    ")
			w.Write(bytes)
			return
		}
		// redirect
		http.Redirect(w, req, "/login", http.StatusOK)
		m["MsgCode"] = 1001
		m["Reason"] = "have not login"
		bytes, _ := json.MarshalIndent(m, "", "    ")
		w.Write(bytes)
		return
	}
	if t.CheckCookie(c.Value) {
		t.M.ServeHTTP(w, req)
		return
	}
	if req.URL.Path == "/login" {
		cookie := &http.Cookie{Name: "mysid", MaxAge: -1, Path: "/"}
		http.SetCookie(w, cookie)
		http.Redirect(w, req, "/login", http.StatusOK)
		m["MsgCode"] = 1002
		m["Reason"] = "session expire"
		bytes, _ := json.MarshalIndent(m, "", "    ")
		w.Write(bytes)
	}
}
func GenSessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("welcome!"))
}

func (t *TUMux) CheckCookie(sid string) bool {
	t.Lock()
	_, ok := t.Sessions[sid]
	t.Unlock()
	return ok
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", login)

	um := &TUMux{
		make(map[string]*TSession),
		mux,
		sync.Mutex{},
	}
	srv := &http.Server{
		Addr:    ":9001",
		Handler: um,
	}
	fmt.Println(srv.ListenAndServe())
}
