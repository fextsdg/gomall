package utils

func MustHandlerError(err error) {
	if err != nil {
		panic(err)
	}
}
