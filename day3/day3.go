package main

import "fmt"

func isOverlapClaim(area []int, fabric [][]int) bool {
	for row := area[1]; row < area[1]+area[3]; row++ {
		for col := area[0]; col < area[0]+area[2]; col++ {
			if fabric[row][col] > 1 {
				return false
			}
		}
	}
	return true
}

func part2() {
	//获取claim后的Fabric
	claimEntites, fabric := part1()

	//循环每个claim
	//计算每个claim在fabric中对应区域的claim次数
	//当对应区域任一地方claim次数大于1则跳出循环
	for _, ce := range claimEntites {
		if isOverlapClaim(ce.area, fabric) {
			fmt.Println(ce.claimId)
		}
	}
}
