package routes

import (
	"net/http"
	"os"

	"github.com/Plaenkler/BatteryHistory/pkg/handler"
)

func ProvideSong(writer http.ResponseWriter, request *http.Request) {
	defer handler.HandlePanic("routes")

	name := request.URL.Query().Get("name")

	fileBytes, err := os.ReadFile("song/"+name+".mp3")
	if err != nil {
		panic(err)
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/octet-stream")
	_, err = writer.Write(fileBytes)
	if err != nil {
		panic(err)
	}
}