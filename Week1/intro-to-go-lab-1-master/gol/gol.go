package main

const alive = 255
const dead = 0

func calculateNextState(p golParams, world [][]byte) [][]byte {
	newWorld := make([][]byte, p.imageHeight)
	for i := range newWorld {
		newWorld[i] = make([]byte, p.imageWidth)
	}
	return world //temp
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	return []cell{}
}
