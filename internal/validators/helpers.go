package validators

func trueValues(v ...bool) int32 {
	var counter int32
	for _, it := range v {
		if it {
			counter++
		}
	}
	return counter
}

func countNotNil(v ...*int) int32 {
	var counter int32
	for _, it := range v {
		if it != nil {
			counter++
		}
	}
	return counter
}
