package language

func PanicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}
