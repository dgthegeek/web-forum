package internals

import (
	"errors"
	"fmt"
	"golang-rest-api-starter/router"
	"log"
	"net/http"
)

type Config struct {
	PORT     string
	Hostname string
}

func (s *Config) Init(router *router.NewRouter) {
	fmt.Println("Server started at ", s.Hostname, s.PORT)
	err := http.ListenAndServe(s.PORT, router)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		log.Fatalf("error starting server: %s\n", err)
	}

}
