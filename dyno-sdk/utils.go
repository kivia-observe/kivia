package dynosdk

func (w *ResponseWriteWrapper) Write(b []byte) (int, error) {
    w.body.Read(b)
    return w.ResponseWriter.Write(b)
}

func (w *ResponseWriteWrapper) WriteHeader(statusCode int) {
    w.statusCode = statusCode
    w.ResponseWriter.WriteHeader(statusCode)
}