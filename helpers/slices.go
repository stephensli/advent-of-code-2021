package helpers

func Clone[T any](input [][]T) [][]T {
	clone := make([][]T, len(input))

	for inputRowIndex, inputRow := range input {
		row := make([]T, len(inputRow))

		for valueIndex, rowValue := range inputRow {
			row[valueIndex] = rowValue
		}

		clone[inputRowIndex] = row
	}

	return clone
}
