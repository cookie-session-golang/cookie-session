package main

import (
	"net/http"
	_ "./cookie"
	"./session"
	"time"
	"html/template"
)

var globalSessions *session.Manager

func main() {
	globalSessions, _  = session.NewManager("memory", "gosessionid", 3600)

	go globalSessions.GC()

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	//cookie.Cookie(w, r)
}

func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)

	r.ParseForm()

	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")

		w.Header().Set("Content-Type", "text/html")

		t.Execute(w, sess.Get("username"))

	} else {
		sess.Set("username", r.Form["username"])

		http.Redirect(w, r, "/", 302)
	}
}

func count(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)

	createtime := sess.Get("createtime")

	if createtime == nil {
		sess.Set("createtime", time.Now().Unix())
	} else if (createtime.(int64) + 360) < (time.Now().Unix()) {
		globalSessions.SessionDestroy(w, r)

		sess = globalSessions.SessionStart(w, r)
	}

	ct := sess.Get("countnum")

	if ct == nil {
		sess.Set("countnum", 1)
	} else {
		sess.Set("countnum", (ct.(int) + 1))
	}

	t, _ := template.ParseFiles("count.gtpl")

	w.Header().Set("Content-Type", "text/html")

	t.Execute(w, sess.Get("countnum"))
}
