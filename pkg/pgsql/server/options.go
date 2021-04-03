package server

import (
	"github.com/codenotary/immudb/pkg/logger"
)

type Option func(s *srv)

func Host(c string) Option {
	return func(args *srv) {
		args.Host = c
	}
}

func Port(port string) Option {
	return func(args *srv) {
		args.Port = port
	}
}

func Logger(logger logger.Logger) Option {
	return func(args *srv) {
		args.Logger = logger
	}
}
