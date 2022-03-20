package Animals

type Lion struct {
}

func (l Lion) MakeNoise() (string, *Error) {
	return ANIMALSOUND, nil
}
