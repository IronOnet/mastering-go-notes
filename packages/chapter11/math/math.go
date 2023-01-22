package math 

func Average(xs []float64) float64{
	total := float64(0) 
	for _, x := range xs{
		total += x
	}
	return total / float64(len(xs))
}

func Min(a, b int) int{
	if a < b{
		return a
	}
	return b 
}

func Max(a, b int) int{
	if a > b{
		return a 
	}
	return b
}