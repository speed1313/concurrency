package main

import "fmt"

// available: the number of available resources of each type
// max: the maximum demand of each process
// allocated: the number of resources of each type currently allocated to each process
// need: the remaining resource need of each process
func isSafe(available []int, allocated [][]int, need [][]int) bool {
	N := len(allocated)
	M := len(available)
	work := make([]int, N)
	copy(work, available)
	finish := make([]bool, N)
	for i := 0; i < N; i++ {
		if finish[i] != false {
			continue
		}
		counter := 0
		for j := 0; j < M; j++ {
			if need[i][j] <= work[j] {
				counter++
			}
		}
		if counter == M {
			for j := 0; j < M; j++ {
				work[j] += allocated[i][j]
			}
			finish[i] = true
			i = -1 // because of this line, this algorithm is O(m*n^2)
			continue
		}
	}
	for i := 0; i < N; i++ {
		if finish[i] == false {
			return false
		}
	}
	return true
}

func checkRequest(request []int, need [][]int, process int, available []int, allocated [][]int) {

	for i := 0; i < len(request); i++ {
		if request[i] <= need[process][i] {
			continue
		} else {
			fmt.Println("Error: request > need")
			return
		}
	}
	for i := 0; i < len(request); i++ {
		if request[i] <= available[i] {
			continue
		} else {
			fmt.Println("Error: request > available")
			return
		}
	}

	for i := 0; i < len(available); i++ {
		available[i] -= request[i]
		allocated[process][i] += request[i]
		need[process][i] -= request[i]
	}
	isSafe := isSafe(available, allocated, need)
	fmt.Println("isSafe: ", isSafe)
	if !isSafe {
		for i := 0; i < len(available); i++ {
			available[i] += request[i]
			allocated[process][i] -= request[i]
			need[process][i] += request[i]
		}
	}

}

func main() {
	available := []int{3, 3, 2}
	max := [][]int{{7, 5, 3}, {3, 2, 2}, {9, 0, 2}, {2, 2, 2}, {4, 3, 3}}
	allocated := [][]int{{0, 1, 0}, {2, 0, 0}, {3, 0, 2}, {2, 1, 1}, {0, 0, 2}}
	need := max
	for i := 0; i < len(max); i++ {
		for j := 0; j < len(max[i]); j++ {
			need[i][j] = max[i][j] - allocated[i][j]
		}
	}

	// request of process 1
	request := []int{1, 0, 2}
	process := 1
	checkRequest(request, need, process, available, allocated)
	// request of process 4
	request = []int{3, 3, 0}
	process = 4
	checkRequest(request, need, process, available, allocated)
	// request of process 0
	request = []int{0, 2, 0}
	process = 0
	checkRequest(request, need, process, available, allocated)

}
