package errors

import "fmt"

// JSONErrors response model
//
// swagger:response JSONErrors
type JSONErrors []JSONError

// JSONError generic response
type JSONError struct {
	Status string            `json:"status"`
	Source map[string]string `json:"source"`
	Title  string            `json:"title"`
	Detail string            `json:"detail"`
}

// New creates new JSONError
func New() JSONErrors {
	return JSONErrors{}
}

// Add adds error to the collection of errors
func (err JSONErrors) Add(status string, source map[string]string, title, detail string) JSONErrors {
	e := JSONError{
		Status: status,
		Source: source,
		Title:  title,
		Detail: detail,
	}
	return append(err, e)
}

// Error formats error
func (je JSONError) Error() string {
	return fmt.Sprintf("%s:%s:%s", je.Status, je.Title, je.Detail)
}
