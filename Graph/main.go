package main

import "fmt"

// Graph Structure
// adjacency list graph
type Graph struct {
	vertices []*Vertex
}

type Vertex struct {
	Key      int
	Adjacent []*Vertex
}

// contains - to check if key is already present

func Contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.Key {
			return true
		}
	}
	return false
}

// Add vertex - a vertex to the graph

func (g *Graph) AddVertex(k int) {
	// to keep uniqueness
	// check if it already exists
	if Contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v not added because it is an existing key", k)
		fmt.Println(err.Error())
		return
	}
	g.vertices = append(g.vertices, &Vertex{Key: k})
}

// Print - to print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, V := range g.vertices {
		fmt.Printf("\n Verted %v : ", V.Key)
		for _, v := range V.Adjacent {
			fmt.Printf("%v", v.Key)
		}
	}
	fmt.Println()
}
func main() {
	test := &Graph{}
	for i := 0; i < 5; i++ {
		test.AddVertex(i)
	}
	//for j := 0; j < 5; j++ {
	//	fmt.Println(test.vertices[j])
	//}
	test.Print()
}
