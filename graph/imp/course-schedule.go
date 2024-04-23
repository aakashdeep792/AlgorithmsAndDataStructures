package graph

/*

207. Course Schedule

There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return true if you can finish all courses. Otherwise, return false.



Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: true
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0. So it is possible.
Example 2:

Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
Output: false
Explanation: There are a total of 2 courses to take.
To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1. So it is impossible.


*/

func buildGraph(node int, edges [][]int) [][]int {
	graph := make([][]int, node)
	for _, v := range edges {
		graph[v[0]] = append(graph[v[0]], v[1])
	}

	return graph
}

func dfsNotCycle(node int, graph [][]int, visited map[int]bool, currStack []bool) bool {
	if visited[node] == true {
		if currStack[node] == true {
			return false //cycle
		}
		return true
	}

	visited[node] = true
	currStack[node] = true
	ans := true
	for _, v := range graph[node] {
		ans = ans && dfsNotCycle(v, graph, visited, currStack)
	}

	currStack[node] = false // back tracking
	return ans
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	visited := make(map[int]bool)
	currStack := make([]bool, numCourses)
	ans := true
	for i := 0; i < len(graph); i++ {
		ans = ans && dfsNotCycle(i, graph, visited, currStack)
	}

	return ans
}
