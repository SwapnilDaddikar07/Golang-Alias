package Animals

type Tiger struct {
}

func (t Tiger) MakeNoise() (string, *Error) {
	return ANIMALSOUND, nil
}
