package main

import "net/http"

func (cfg *apiConfig) handlerReset(writer http.ResponseWriter, request *http.Request) {
	cfg.fileserverHits = 0
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Hits reset to 0"))
}
