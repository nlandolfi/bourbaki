graph.csv: ../sheets/*/*.tex
	bbk_implicit_graph > graph.csv
graph.graphviz: graph.csv graph.go *.tmpl
	go run graph.go --tmpl graph.tmpl > graph.graphviz
force.csv: graph.csv graph.go *.tmpl
	go run graph.go --tmpl force.tmpl > force.csv
graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf
o: graph.pdf
	open graph.pdf
check:
	go run check.go
clean:
	rm graph.csv
