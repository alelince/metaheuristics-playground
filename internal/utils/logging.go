package utils

import "github.com/rs/zerolog"

func InitLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
