package apperror

import "github.com/rs/zerolog"

func LogError(logger zerolog.Logger, source string, err error) {
	if appErr, ok := err.(*AppError); ok {
		event := logger.Error().Str("source", source)
		for k, v := range appErr.LogFields() {
			event = event.Interface(k, v)
		}
		event.Msg("AppError occurred")
	} else {
		logger.Error().
			Str("source", source).
			Err(err).
			Msg("Unexpected error")
	}
}
