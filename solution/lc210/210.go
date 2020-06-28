package lc210

func findZeroDependency(dependencyCounts []int, visited []bool) int {
	for i, count := range dependencyCounts {
		if count == 0 && !visited[i] {
			return i
		}
	}
	return -1
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	dependencyCounts := make([]int, numCourses)
	visited := make([]bool, numCourses)
	result := make([]int, 0)
	paths := make(map[int][]int)
	for _, edge := range prerequisites {
		current, prerequist := edge[0], edge[1]
		links, ok := paths[prerequist]
		if !ok {
			links = make([]int, 0)
		}
		links = append(links, current)
		paths[prerequist] = links
		dependencyCounts[current]++
	}
	for len(result) != numCourses {
		idx := findZeroDependency(dependencyCounts, visited)
		if idx == -1 {
			return make([]int, 0)
		}
		visited[idx] = true
		result = append(result, idx)
		links, ok := paths[idx]
		if ok {
			for _, nextIdx := range links {
				dependencyCounts[nextIdx]--
			}
		}
	}

	return result
}
