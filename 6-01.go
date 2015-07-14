/*
 * DAILY PROGRAMMER - 6/01/2015
 * LUMBERJACK ALOGRITHM
 * http://www.reddit.com/r/dailyprogrammer/comments/378h44/20150525_challenge_216_easy_texas_hold_em_1_of_3/
 * Author: Muhan Guclu
 * Description: Calculates most even distribution of "logs" amongst existing uneven "piles"
 */

package main

import "fmt"

func getSize () int {
    fmt.Println("Grid size?")

    var size int
    _, err := fmt.Scanf("%d", &size)
    if err == nil {
        return size
    } else {
        return 0
    }
}

func getLogs () int {
    fmt.Println("Number of logs to place?")

    var logs int
    _, err := fmt.Scanf("%d", &logs)
    if err == nil {
        return logs
    } else {
        return 0
    }
}

func getLogsInRow (size int, row int) []int {
    fmt.Println("Enter value of each pile in ROW", row+1)

    ret := make([]int, size)
    for i := range ret {
        var val int
        fmt.Scanf("%d", &val)
        ret[i] = val
    }

    return ret
}

func getGrid () ([][]int, int) {
    size := getSize()
    logs := getLogs()
    grid := make([][]int, size)

    for rowIdx := range grid {
        row := getLogsInRow(size, rowIdx)
        grid[rowIdx] = row
    }

    return grid, logs
}

func addLogs (grid [][]int, logs int) [][]int {
    for logs > 0 {
        grid = incrementMinimum(grid)
        logs--
    }
    return grid
}

func contains(grid [][]int, b int) bool {
    for i := range grid {
        for j := range grid {
            if grid[i][j] == b {
                return true 
            }
        }
    }
    return false
}

func minimum (grid [][]int) int {
    ret := grid[0][0]

    for i := range grid {
        for j := range grid {
            if grid[i][j] < ret {
                ret = grid[i][j]
            }
        }
    }

    return ret
}

func incrementMinimum (grid [][]int) [][]int {
    min := minimum(grid)
    hasIncremented := false
    for i := range grid {
        for j := range grid {
            if !hasIncremented && grid[i][j] == min {
                grid[i][j]++
                hasIncremented = true
            }
        }
    }
    return grid
}

func main () {
    grid, logs := getGrid()
    grid = addLogs(grid, logs)
    fmt.Println("Result:", grid)
}
