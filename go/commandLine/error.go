package main

type RequestError struct {
	HTTPCore int
	Body     string
	Err      string
}

func (r RequestError) Error() string {
	return r.Err
}
