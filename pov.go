package pov

const testVersion = 2

type Graph struct {
	arcPairs []edge
	root     string
}

type edge struct{ fr, to string }

func New() *Graph {
	return new(Graph)
}
func (g *Graph) AddNode(nodeLabel string) {
}
func (g *Graph) AddArc(from, to string) {
	g.arcPairs = append(g.arcPairs, edge{from, to})
}
func (g *Graph) ArcList() (o []string) {
	for _, a := range g.arcPairs {
		o = append(o, formatArc(a))
	}
	return
}
func formatArc(a edge) string {
	return a.fr + " -> " + a.to
}
func (g *Graph) ChangeRoot(oldRoot, newRoot string) *Graph {
	fr, to := oldRoot, newRoot
	for to != fr {
		for i := 0; i < len(g.arcPairs); i++ {
			if g.arcPairs[i].to == to {
				to = g.arcPairs[i].fr
				g.arcPairs[i].fr, g.arcPairs[i].to = g.arcPairs[i].to, g.arcPairs[i].fr
			}
		}
	}

	return g
}
