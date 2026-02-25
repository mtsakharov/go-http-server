package httpcore

type StatusCode int

const (
	StatusOK                  StatusCode = 200
	StatusCreated             StatusCode = 201
	StatusNotFound            StatusCode = 404
	StatusInternalServerError StatusCode = 500
)

var statusText = map[StatusCode]string{
	StatusOK:                  "OK",
	StatusCreated:             "Created",
	StatusNotFound:            "Not Found",
	StatusInternalServerError: "Internal Server Error",
}
