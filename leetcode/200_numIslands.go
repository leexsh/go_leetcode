package leetcode

/*
題目：岛屿的数量
给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
*/
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				dfslands(grid, i, j)
				res++
			}
		}
	}
	return res
}

func dfslands(grid [][]byte, i, j int) {
	w, h := len(grid[0]), len(grid)
	if i < 0 || j < 0 || i >= h || j >= w || grid[i][j] != '1' {
		return
	}
	grid[i][j] = '2'
	dfslands(grid, i-1, j)
	dfslands(grid, i+1, j)
	dfslands(grid, i, j-1)
	dfslands(grid, i, j+1)

}
