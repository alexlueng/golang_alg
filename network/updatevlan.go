package network

// Feature 5 Efficiently update the VLAN IDs of network switches in a gird topology

func dfsVLAN(matrix *[][]int, r, c, currId, newId int) {
	if r < 0 || c < 0 || r >= len((*matrix)) || c >= len((*matrix)[0]) || (*matrix)[r][c] != currId {
		return
	}

	(*matrix)[r][c] = newId

	dfsVLAN(matrix, r-1, c, currId, newId)
	dfsVLAN(matrix, r+1, c, currId, newId)
	dfsVLAN(matrix, r, c-1, currId, newId)
	dfsVLAN(matrix, r, c+1, currId, newId)
}

func updateVLAN(matrix [][]int, r, c, newId int) [][]int {
	currId := matrix[r][c]
	if currId == newId {
		return matrix
	}
	dfsVLAN(&matrix, r, c, currId, newId)
	return matrix
}
