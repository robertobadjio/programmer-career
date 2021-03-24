package part4

import "fmt"

// Сложность по времени O(|V|+|E|)
func ExistsRoute() {
	graph := New()
	node0 := graph.AddNode()
	node1 := graph.AddNode()
	node2 := graph.AddNode()
	node3 := graph.AddNode()
	node4 := graph.AddNode()

	graph.AddEdge(node0, node1, 0)
	graph.AddEdge(node0, node2, 0)
	graph.AddEdge(node1, node3, 0)
	graph.AddEdge(node1, node4, 0)

	// 0 -> 3
	var visited [10]bool
	fmt.Printf("Модификация метода обхода в глубину - %t\n", search(0, 3, visited, graph))
}

// visited — Массив цветов вершин
// t — Конечная вершина
func search(u, t int, visited [10]bool, graph *Graph) bool {
	if u == t {
		return true
	}

	visited[u] = true // Помечаем вершину как пройденную

	for _, v := range graph.Neighbors(u) {
		if !visited[v] {
			if search(v, t, visited, graph) {
				return true
			}
		}
	}

	return false
}
