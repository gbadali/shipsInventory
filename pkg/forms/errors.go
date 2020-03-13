package forms

type errors map[string][]string

// Add takes in the field (what html form field the message goes to)
// and a message (the error to show to the user) and adds them to the
// error type
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Get retrieves an error message from the error type
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
