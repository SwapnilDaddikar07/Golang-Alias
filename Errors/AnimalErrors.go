package Errors

type AnimalErrors struct {
	ErrorCode    int
	ErrorMessage string
}

func (e AnimalErrors) SetError() {
	e.ErrorMessage = "error message"
}

//Add new methods wheneever needed