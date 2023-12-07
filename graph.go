package main

import (
	"io"
	"log"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Graph struct {
	Nodes []opts.GraphNode
	Links []opts.GraphLink

	Graph *charts.Graph
}

func NewGraph() Graph {
	var graph Graph
	graph.Nodes = make([]opts.GraphNode, 0)
	graph.Links = make([]opts.GraphLink, 0)
	return graph
}

func GenerateGraph(mods map[string]Mod) Graph {
	var graph Graph = NewGraph()

	for _, mod := range mods {
		node := opts.GraphNode{
			Name: mod.Id,
		}
		graph.Nodes = append(graph.Nodes, node)

		for k := range mod.DepsMap {
			if _, ok := mods[k]; ok {
				graph.Links = append(graph.Links, opts.GraphLink{Source: node.Name, Target: k})
			}
		}
	}

	graph.Graph = charts.NewGraph()
	graph.Graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Dependency Graph"}),
	)
	graph.Graph.AddSeries("graph", graph.Nodes, graph.Links).
		SetSeriesOptions(
			charts.WithGraphChartOpts(opts.GraphChart{
				Force:              &opts.GraphForce{Repulsion: 8000},
				Layout:             "force",
				Roam:               true,
				FocusNodeAdjacency: true,
				EdgeSymbol:         []string{"circle", "arrow"},
				EdgeSymbolSize:     []int{4, 10},
			}),
			charts.WithEmphasisOpts(opts.Emphasis{
				Label: &opts.Label{
					Show:     true,
					Color:    "black",
					Position: "inside",
				},
			}),
			// charts.WithLineStyleOpts(opts.LineStyle{
			// 	Curveness: 0.3,
			// }),
		)

	return graph
}

func (graph Graph) Render() error {
	page := components.NewPage()
	page.AddCharts(graph.Graph)

	f, err := os.Create("html/graph.html")
	if err != nil {
		return err
	}
	page.Render(io.MultiWriter(f))

	log.Println("Generated graph in 'html/graph.html'")
	return nil
}
