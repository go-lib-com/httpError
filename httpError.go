package httpError


type Err struct {
    Code    int
    Key     string
    Message string
    Err error
}

func New(code int, key string, msg string, err error) *Err {
    return &Err{
        Code:    code,
        Key:     key,
        Message: msg,
        Err: err,
    }
}

// Error makes it compatible with `error` interface.
func (e *Err) Error() string {
    return e.Key + ": " + e.Message
}

