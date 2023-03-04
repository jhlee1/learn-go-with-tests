package iteration

func Repeat(s string, repeat int) (output string) {
	for i := 0; i < repeat; i++ {
		output += s
	}

	return
}
