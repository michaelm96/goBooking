package forms

type errors map[string][]string

// Add error message for given field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get returns the first error message
func (e errors) Get(field string) string {
	errStr := e[field]
	if len(errStr) == 0 {
		return ""
	}

	return errStr[0]
}
