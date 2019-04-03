package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	/**
	读取文件信息
	*/
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col) //这里读取文件信息 给row 和col赋值

	maze := make([][]int, row) //创建二维切片
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

/**
开始走迷宫
*/
type point struct {
	i, j int
}

var dirs = [4]point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}} //上 左 下 右

func (p point) add(r point) point { //探索
	return point{p.i + r.i, p.j + r.j}
}
func (p point) at(grid [][]int) (int, bool) { //校验当前位置是否能走
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}
func walk(maze [][]int, start, end point) [][]int { //start 是起点,end是终点
	steps := make([][]int, len(maze)) //大小是len(maze)行的切片
	for i := range steps {            //i是切片的行位置
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start} //创建一个队列 类型是point,值为start

	for len(Q) > 0 { //队列不空继续走
		cur := Q[0]
		Q = Q[1:]

		if cur == end { //到达终点退出
			break
		}
		for _, dir := range dirs {
			next := dir.add(cur) //返回的是坐标
			val, ok := next.at(maze)
			if !ok || val == 1 { //越界或者撞墙了
				continue //跳出本次循环
			}
			val, ok = next.at(steps)
			if !ok || val != 0 { //越界走过了
				continue
			}
			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)         //返回的是当前位置的值
			steps[next.i][next.j] = curSteps + 1 //设此位置值加1

			Q = append(Q, next)
		}
	}
	return steps
}

func main() {
	maze := readMaze("maze/maze.in")
	/*for _, row := range maze {//测试读取地图是否正确
		for _,val := range row{
			fmt.Printf("%d ",val)
		}
		fmt.Println()
	}*/
	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val) //3d 3位对齐
		}
		fmt.Println()
	}
}
