package main

import (
	"flag"
	"log/slog"
	"os"

	userv1 "solabar-server/api/user/v1"
	userservice "solabar-server/internal/user/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/kanengo/ktp/transport/hertz"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func newServer(userSer *userservice.UserExtService) *hertz.Server {
	svr := hertz.NewServer()

	userv1.RegisterUserExtHertzServer(svr, userSer, nil)

	return svr
}

func newApp(hs *hertz.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Server(hs),
	)
}

func main() {
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		// AddSource: true,
		Level: slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.SourceKey {
				return slog.Attr{}
			}
			return a
		},
	}))
	app, err := wireApp(logger)
	if err != nil {
		panic(err)
	}
	if err := app.Run(); err != nil {
		panic(err)
	}
}
