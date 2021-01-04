graph.graphviz: graph.csv graph.go *.tmpl
	go run graph.go --tmpl graph.tmpl > graph.graphviz
force.csv: graph.csv graph.go *.tmpl
	go run graph.go --tmpl force.tmpl > force.csv
graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf
o: graph.pdf
	go run check.go
	open graph.pdf
