graph.csv: ../sheets/*/*.tex
	bbk_implicit_graph > graph.csv
graph.graphviz: graph.csv *.tmpl
	bbk_graph --tmpl graph.tmpl > graph.graphviz
graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf
force.csv: graph.csv *.tmpl
	go run graph.go --tmpl force.tmpl > force.csv
o: graph.pdf
	open graph.pdf
check:
	bbk_check
clip:
	bbk_clip
clean:
	rm graph.csv
