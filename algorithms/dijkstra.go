package algorithms

import (
	"fmt"
	"math"
)

/*
	Dijkstra's Algorithm:
	- Dijkstra's algorithm initializing dist[s] to 0 and all other distTo[] entries to positive infinity.
	  Then, it repeatedly relaxes and adds to the tree a non-tree vertex with the lowest distTo[] value,
	  continuing until all vertices are on the tree or no non-tree vertex has a finite distTo[] value.
*/

// Define an infinite value for initial distances
const INF = math.MaxInt64

// Graph represents a graph with a map of its adjacency list
type Graph struct {
	vertices map[int][]Edge
}

// Edge represents an edge in the graph
type Edge struct {
	to     int
	weight int
}

// NewGraph creates a new Graph
func NewGraph() *Graph {
	return &Graph{vertices: make(map[int][]Edge)}
}

// AddEdge adds an edge to the graph
func (g *Graph) AddEdge(from, to, weight int) {
	g.vertices[from] = append(g.vertices[from], Edge{to, weight})
}

// Dijkstra finds the shortest path using Dijkstra's algorithm
func (g *Graph) Dijkstra(start int) (dist map[int]int, prev map[int]int) {
	// Initialize distances and previous vertices
	dist = make(map[int]int)
	prev = make(map[int]int)
	for v := range g.vertices {
		dist[v] = INF
		prev[v] = -1
	}
	dist[start] = 0

	// Create a priority queue
	pq := make(PriorityQueue, 0)
	pq.Push(&Item{value: start, priority: 0})

	// Main loop of the algorithm
	for len(pq) > 0 {
		u := pq.Pop().value
		for _, edge := range g.vertices[u] {
			v := edge.to
			alt := dist[u] + edge.weight
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				pq.Push(&Item{value: v, priority: alt})
			}
		}
	}
	return dist, prev
}

// PriorityQueue and related methods for the priority queue
type PriorityQueue []*Item

type Item struct {
	value    int
	priority int
	index    int
}

func (pq *PriorityQueue) Push(x *Item) {
	n := len(*pq)
	x.index = n
	*pq = append(*pq, x)
	pq.up(n)
}

func (pq *PriorityQueue) Pop() *Item {
	old := *pq
	n := len(old)
	item := old[0]
	old[0] = old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	pq.down(0)
	return item
}

func (pq *PriorityQueue) up(j int) {
	for {
		i := (j - 1) / 2
		if i == j || (*pq)[i].priority <= (*pq)[j].priority {
			break
		}
		pq.swap(i, j)
		j = i
	}
}

func (pq *PriorityQueue) down(i int) {
	n := len(*pq)
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && (*pq)[j2].priority < (*pq)[j1].priority {
			j = j2
		}
		if (*pq)[j].priority >= (*pq)[i].priority {
			break
		}
		pq.swap(i, j)
		i = j
	}
}

func (pq *PriorityQueue) swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
	(*pq)[i].index, (*pq)[j].index = i, j
}

func TestDijkstraAlgorithm() {
	graph := NewGraph()
	// Example: Add edges to the graph
	graph.AddEdge(0, 1, 4)
	graph.AddEdge(0, 2, 2)
	graph.AddEdge(1, 2, 5)
	graph.AddEdge(1, 3, 10)
	graph.AddEdge(2, 3, 3)
	graph.AddEdge(3, 4, 1)

	dist, prev := graph.Dijkstra(0)
	fmt.Println("Distance:", dist)
	fmt.Println("Previous:", prev)
}
