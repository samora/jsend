package jsend

// H is hash containing any arbitrary values
type H map[string]interface{}

// Success JSend response
func Success(data interface{}) H {
	return H{
		"status": "success",
		"data":   data,
	}
}

// Fail JSend response
func Fail(data interface{}) H {
	return H{
		"status": "fail",
		"data":   data,
	}
}

// Error JSend response
func Error(msg string, code int64, data interface{}) H {
	return H{
		"status":  "error",
		"message": msg,
		"code":    code,
		"data":    data,
	}
}
