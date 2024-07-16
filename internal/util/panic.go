package util

func PanicIF(err error) {
	if err != nil {
		panic(err)
	}
}
