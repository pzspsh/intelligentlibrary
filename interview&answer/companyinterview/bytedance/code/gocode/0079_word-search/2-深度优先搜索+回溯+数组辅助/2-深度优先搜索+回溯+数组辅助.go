package main

import "fmt"

func main() {
	board := [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	fmt.Println(exist1(board, "abcd"))
}

// leetcode79_单词搜索
func exist1(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs1(board, i, j, word, 0, visited) {
				return true
			}
		}
	}
	return false
}

func dfs1(board [][]byte, i, j int, word string, level int, visited [][]bool) bool {
	res := false
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) &&
		visited[i][j] == false && board[i][j] == word[level] {
		if level == len(word)-1 {
			return true
		}
		visited[i][j] = true
		level = level + 1
		res = dfs1(board, i+1, j, word, level, visited) ||
			dfs1(board, i-1, j, word, level, visited) ||
			dfs1(board, i, j+1, word, level, visited) ||
			dfs1(board, i, j-1, word, level, visited)
		if !res {
			visited[i][j] = false
			level = level - 1
		}
	}
	return res
}
