package handlers

import (
	"net/http"
	"github.com/go-redis/redis"
)

func GogetHandler(db *redis.Client) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		uuid := r.URL.Path
		val := db.Get("uuid:" + uuid).Val()
		if(val != "") {
			http.Redirect(w, r, val, http.StatusSeeOther)
			return
		}
		http.NotFound(w, r)
	}
	return http.HandlerFunc(handler)
}