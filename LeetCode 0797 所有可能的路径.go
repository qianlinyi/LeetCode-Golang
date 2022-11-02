package main

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	pre, path, vis, ans := make([]int, n), make([]int, n), make([]bool, n), make([][]int, 0)
	var dfs func(v int)
	var findPath func(v int)
	findPath = func(v int) {
		if v == 0 {
			path = append(path, v)
			return
		}
		findPath(pre[v])
		path = append(path, v)
	}
	dfs = func(v int) {
		if v == n-1 {
			path = []int{}
			findPath(v)
			ans = append(ans, path)
			return
		}
		for _, u := range graph[v] {
			if !vis[u] {
				vis[u] = true
				pre[u] = v
				dfs(u)
				vis[u] = false
			}
		}
	}
	dfs(0)
	return ans
}
