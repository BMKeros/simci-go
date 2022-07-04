package bootstrap

import "simci-go/internal/platform/server"

func Run() error {
	srv := server.New("host", 8080)

	return srv.Run()
}
