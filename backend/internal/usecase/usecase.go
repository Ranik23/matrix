package usecase

import (
	"ProjMatrix/internal/storage/postgres"
	"ProjMatrix/internal/storage/redis"
	"ProjMatrix/pkg/wpool"
	"errors"
	"fmt"
	"sync"
	"time"
	mtrx "ProjMatrix/pkg/matrix"
)





type UseCase interface {
	CalculateLinearForm(matrices [][][]float64, coefficients []float64) ([][]float64, float64, error)
    CalculateParallelLinearForm(matrices [][][]float64, coefficients []float64) ([][]float64, float64, error)
    CalculatePolynomial(matrix, identityMatrix [][]float64, coefficients []float64) ([][]float64, float64, error)
	CalculateParallelPolynomial(matrix, identityMatrix [][]float64, coefficients []float64) ([][]float64, float64, error)
}

type UserOperator struct {
	poolWorker		*wpool.WorkerPool
	postgresClient 	postgres.Storage
	redisCient     	redis.CacheStorage
}


func (uc *UserOperator) CalculateLinearForm(matrices [][][]float64, coefficients []float64) ([][]float64, float64, error) {
	start := time.Now()

	if len(matrices) == 0 || len(coefficients) == 0 {
		return nil, 0, errors.New("the array of matrices or the array of coefficients is empty")
	}

	if len(matrices) != len(coefficients) {
		return nil, 0, errors.New("the number of matrices and coefficients must match")
	}

	rows, cols := len(matrices[0]), len(matrices[0][0])

	result := make([][]float64, rows)
	for i := range result {
		result[i] = make([]float64, cols)
	}

	for k, matrix := range matrices {
		coefficient := coefficients[k]
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				result[i][j] += coefficient * matrix[i][j]
			}
		}
	}

	elapsed := time.Since(start).Seconds()

	return result, elapsed, nil
}

func (uc *UserOperator) CalculateParallelLinearForm(matrices [][][]float64, coefficients []float64) ([][]float64, float64, error) {
	start := time.Now()
	if len(matrices) == 0 || len(coefficients) == 0 {
		return nil, 0, errors.New("the array of matrices or the array of coefficients is empty")
	}

	if len(matrices) != len(coefficients) {
		return nil, 0, errors.New("the number of matrices and coefficients must match")
	}

	rows, cols := len(matrices[0]), len(matrices[0][0])

	result := make([][]float64, rows)
	for i := range result {
		result[i] = make([]float64, cols)
	}

	var mu sync.Mutex

	for k, matrix := range matrices {
		coefficient := coefficients[k]

		// передаем задачу в пул воркеров
		uc.poolWorker.Submit(func() {
			partialResult := make([][]float64, rows)
			for i := range partialResult {
				partialResult[i] = make([]float64, cols)
			}

			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					partialResult[i][j] = coefficient * matrix[i][j]
				}
			}

			mu.Lock()
			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					result[i][j] += partialResult[i][j]
				}
			}
			mu.Unlock()
		})
	}

	uc.poolWorker.Wait()

	elapsed := time.Since(start).Seconds()

	return result, elapsed, nil
}

func (uc *UserOperator) CalculatePolynomial(matrix, identityMatrix [][]float64, coefficients []float64) ([][]float64, float64, error) {
	start := time.Now()

	if len(matrix) == 0 || len(identityMatrix) == 0 {
		return nil, 0, errors.New("the matrix or the unit matrix is empty")
	}
	if len(matrix) != len(matrix[0]) {
		return nil, 0, errors.New("the matrix must be square")
	}

	n := len(matrix)

	result := make([][]float64, n)
	for i := range result {
		result[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] += coefficients[0] * identityMatrix[i][j]
		}
	}

	// Временная матрица для хранения степеней матрицы A
	currentPower := identityMatrix
	var err error

	for m := 1; m < len(coefficients); m++ {

		currentPower, err = mtrx.MatrixMultiply(currentPower, matrix)
		if err != nil {
			return nil, 0, fmt.Errorf("the Matrix Polynomial could not be calculated: %w", err)
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				result[i][j] += coefficients[m] * currentPower[i][j]
			}
		}
	}

	elapsed := time.Since(start).Seconds()
	return result, elapsed, nil
}

func (uc *UserOperator) CalculateParallelPolynomial(matrix, identityMatrix [][]float64, coefficients []float64) ([][]float64, float64, error) {
	start := time.Now()

	if len(matrix) == 0 || len(identityMatrix) == 0 {
		return nil, 0, errors.New("матрица или единичная матрица пустые")
	}
	if len(matrix) != len(matrix[0]) {
		return nil, 0, errors.New("матрица должна быть квадратной")
	}

	n := len(matrix)

	result := make([][]float64, n)
	for i := range result {
		result[i] = make([]float64, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] += coefficients[0] * identityMatrix[i][j]
		}
	}

	currentPower := identityMatrix
	var err error

	for m := 1; m < len(coefficients); m++ {
		currentPower, err = uc.ParallelMatrixMultiply(currentPower, matrix)
		if err != nil {
			return nil, 0, fmt.Errorf("не удалось вычислить Матричный Полином: %w", err)
		}

		uc.poolWorker.Submit(func() {
			for i := 0; i < n; i++ {
				for j := 0; j < n; j++ {
					result[i][j] += coefficients[m] * currentPower[i][j]
				}
			}
		})
	}

	uc.poolWorker.Wait()
	elapsed := time.Since(start).Seconds()
	return result, elapsed, nil
}

func (uc *UserOperator) ParallelMatrixMultiply(a, b [][]float64) ([][]float64, error) {
	if len(a) == 0 || len(b) == 0 || len(a[0]) != len(b) {
		return nil, errors.New("matrices are incompatible for multiplication")
	}

	rowsA, colsA := len(a), len(a[0])
	colsB := len(b[0])

	result := make([][]float64, rowsA)
	for i := range result {
		result[i] = make([]float64, colsB)
	}

	var mu sync.Mutex
	for i := 0; i < rowsA; i++ {
		row := i
		uc.poolWorker.Submit(func() {
			for j := 0; j < colsB; j++ {
				sum := 0.0
				for k := 0; k < colsA; k++ {
					sum += a[row][k] * b[k][j]
				}
				mu.Lock()
				result[row][j] = sum
				mu.Unlock()
			}
		})
	}

	uc.poolWorker.Wait()
	return result, nil
}