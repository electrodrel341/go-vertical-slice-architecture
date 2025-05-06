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

var ErrorSetDefaultAIModelForProvider = ErrorData{
	CodeValue:       1004,
	CodeDescription: "ERROR_NOT_SET_DEFAULT_AI_MODEL_FOR_PROVIDER",
	Message:         "Не задано значение AI модели по умолчанию для провайдера",
}

var ErrorWrongAIModelForProvider = ErrorData{
	CodeValue:       1005,
	CodeDescription: "ERROR_WRONG_AI_MODEL_FOR_PROVIDER",
	Message:         "Значение AI модели не соответствует провайдеру",
}
