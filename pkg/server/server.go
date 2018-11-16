package server

import (
	"github.com/patrickmn/go-cache"
	"github.com/gorilla/mux"
	"net/http"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

type Server struct {
	urlCache *cache.Cache
}

func New() *Server {

	uCache := cache.New(10, 5)
	s := &Server{
		urlCache: uCache,
	}
	router := mux.NewRouter()
	router.HandleFunc("/url/{longUrl}", s.shortenUrl).Methods("POST")
	http.ListenAndServe(":8080", router)
	return s
}

func (s *Server) shortenUrl(res http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	longUrl := vars["longUrl"]
	md5Hash := s.getMD5Hash(longUrl)
	shortUrl := md5Hash[0:6]
	s.urlCache.Add(shortUrl, longUrl, 100)
	fmt.Printf("long url : %s, short url : %x\n", longUrl, shortUrl)
	res.Write([]byte(shortUrl))
}

func (s *Server) getMD5Hash( longUrl string) string {
	hasher := md5.New()
	hasher.Write([]byte(longUrl))
	return hex.EncodeToString(hasher.Sum(nil))
}

