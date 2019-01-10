package handlers

import (
	"net/http"
	"math/rand"
	"github.com/go-redis/redis"
	"strings"
)

const letterBytes = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randUUID(n int) string {
    b := make([]byte, n)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func getValidUUID(db *redis.Client) string {
	uuid := randUUID(8)
	for db.Exists("uuid:" + uuid).Val() != 0 {
		uuid = randUUID(8)
	}
	return uuid
}

func GomakeHandler(db *redis.Client) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		if val := db.Exists("url:" + url).Val(); val == 0 {			
			uuid := getValidUUID(db)
			db.Set("uuid:" + uuid, url, 0)
			db.Set("url:" + url, uuid, 0)
		}
		uuid := db.Get("url:" + url).Val()
		redirectURL := "http://" + r.Host + "?uuid=" + uuid
		http.Redirect(w, r, redirectURL, http.StatusSeeOther)
	}
	return http.HandlerFunc(handler)
}