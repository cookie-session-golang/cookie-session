package cookie

import (
	"net/http"
	"time"
)

func Cookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now()

	cookie := http.Cookie {
				Name: "username",
				Value: "astaxie",
				Expires: expiration,
	}

	 http.SetCookie(w, &cookie)
}
