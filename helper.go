package adventhelperfunctions

func MapToImage[T any](m map[[2]int]T) [][]T {
	minx, miny := 0, 0
	maxx, maxy := 0, 0
	for key := range m {
		if key[0] < minx {
			minx = key[0]
		}
		if key[1] < miny {
			miny = key[1]
		}
		if key[0] > maxx {
			maxx = key[0]
		}
		if key[1] > maxy {
			maxy = key[1]
		}
	}
	img := make([][]T, maxx-minx)
	for i := range img {
		img[i] = make([]T, maxy-miny)
		for coords, value := range m {
			img[coords[0]-minx][coords[1]-miny] = value
		}
	}
	return img

}
