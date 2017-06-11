package main

import "golang.org/x/tour/pic"

func SwirlPic(dx, dy int) [][]uint8 {
        matrix := make([][]uint8, dy)
        for i := range matrix {
                matrix[i] = make([]uint8, dx)
                for j := range matrix[i] {
                        matrix[i][j] = uint8(j * i - i - j)
                }
        }
        return matrix
}

func GradientPic(dx, dy int) [][]uint8 {
		matrix := make([][]uint8, dy)
		for i := range matrix {
				matrix[i] = make([]uint8, dx)
				for j := range matrix[i] {
						matrix[i][j] = uint8(i * dx + j)
				}
		}
		return matrix
}

func ShowPic() {
		pic.Show(GradientPic)
}


