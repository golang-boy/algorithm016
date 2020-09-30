package Week_03

// 判断当前位置是否能放皇后，（检测规则，记住以前放置的位置，拿新放的进行检测，看新的能不能放）
// 一行放完进行下一行放置
// 下一行如果没法放置，则回到上一行重新放置
// 要放n个

//

type board struct {
	// 竖线，左斜对角线，右斜对角线皇后状态
	columns         map[int]bool
	diagonals_left  map[int]bool
	diagonals_right map[int]bool

	// 棋盘内容
	content [][]byte

	// 皇后数,列数
	n int
}

var ans = [][]string{}

func (b *board) setup(n int) {
	// 竖线，左斜对角线，右斜对角线，
	b.columns, b.diagonals_left, b.diagonals_right = map[int]bool{}, map[int]bool{}, map[int]bool{}

	b.n = n
	b.content = make([][]byte, n)

	for r := 0; r < n; r++ {
		b.content[r] = make([]byte, n)
		for l := 0; l < n; l++ {
			b.content[r][l] = '.'
		}
	}
}

func (b *board) isValid(row, col int) bool {
	if b.columns[col] || b.diagonals_left[col-row] || b.diagonals_right[col+row] {
		return false
	}
	return true
}

func (b *board) setQueen(row, col int) {
	b.content[row][col] = 'Q'
	b.columns[col] = true
	b.diagonals_left[col-row] = true
	b.diagonals_right[col+row] = true
}

func (b *board) unSetQueen(row, col int) {
	b.content[row][col] = '.'
	delete(b.columns, (col))
	delete(b.diagonals_left, (col - row))
	delete(b.diagonals_right, (col + row))
}

func solveNQueens(n int) [][]string {
	var b = new(board)
	b.setup(n)
	backtrack(b, 0)
	return ans
}

func backtrack(b *board, row int) {
	if row == b.n {
		//最后一行的皇后也处理完了
		res := []string{}
		for _, c := range b.content {
			res = append(res, string(c))
		}
		ans = append(ans, res)
		return
	}

	// 当前行中有n列, 依次选择这n列，进行判断
	for col := 0; col < b.n; col++ {
		if !b.isValid(row, col) {
			continue
		}

		b.setQueen(row, col)
		backtrack(b, row+1)
		b.unSetQueen(row, col)
	}
}
