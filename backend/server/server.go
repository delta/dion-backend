package server

import (
	"fmt"
	"net/http"
	"time"

	"delta.nitt.edu/dion/config"
	"delta.nitt.edu/dion/router"
	"delta.nitt.edu/dion/services/logging"
)

func StartServer() {
	router := router.InitRouter()
	maxHeaderBytes := 1 << 20
	logging.Sugared().Infof("Starting the server and listening on port %d", config.C.Server.Port)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", config.C.Server.Port),
		Handler:        router,
		ReadTimeout:    (time.Duration(config.C.Server.ReadTimeout)) * (time.Second),
		WriteTimeout:   (time.Duration(config.C.Server.WriteTimeout)) * (time.Second),
		MaxHeaderBytes: maxHeaderBytes,
		// ErrorLog: (*log.Logger)(logging.L.Named("gin error:")),
	}
	logging.Sugared().Infof("Starting the server and listening on port %d", config.C.Server.Port)
	s.ListenAndServe()
}
