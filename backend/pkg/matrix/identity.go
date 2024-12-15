package matrix

func GenerateIdentityMatrix(size int) [][]float64 {
	if size <= 0 {
		return nil // Возвращаем nil, если размер некорректен
	}

	identityMatrix := make([][]float64, size)
	for i := 0; i < size; i++ {
		identityMatrix[i] = make([]float64, size)
		identityMatrix[i][i] = 1
	}

	return identityMatrix
}
