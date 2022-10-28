package utils

func ResponseFailed(msg string) map[string]interface{} {
    return map[string]interface{}{
        "status":  "error",
        "message": msg,
    }
}

func ResponseSuccessNoData(msg string) map[string]interface{} {
    return map[string]interface{}{
        "status":  "Success",
        "message": msg,
    }
}

func ResponseSuccessWithData(msg string, data interface{}) map[string]interface{} {
    return map[string]interface{}{
        "status":  "Success",
        "message": msg,
        "data":    data,
    }
}
