package apperror

type ErrorData struct {
	CodeValue       int
	CodeDescription string
	Message         string
}

var ErrorInternalServerError = ErrorData{
	CodeValue:       5001,
	CodeDescription: "INTERNAL_SERVER_ERROR",
	Message:         "Внутренняя, обратитесь в техподдержку",
}

var ErrorNotFoundOpenAiApiKey = ErrorData{
	CodeValue:       1001,
	CodeDescription: "ERROR_NOT_FOUND_OPENAI_API_KEY",
	Message:         "Ошибка получения конфига, обратитесь в техподдержку",
}

var ErrorWrongAIApiProvider = ErrorData{
	CodeValue:       1002,
	CodeDescription: "ERROR_WRONG_AI_PROVIDER",
	Message:         "Некорректное значение AI провайдера",
}

var ErrorWrongAIModel = ErrorData{
	CodeValue:       1003,
	CodeDescription: "ERROR_WRONG_AI_MODEL",
	Message:         "Некорректное значение AI модели",
}
