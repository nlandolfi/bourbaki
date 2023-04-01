graph.csv: ../sheets/*/*.lit
	bbk_implicit_graph > graph.csv
graph.graphviz: graph.csv *.tmpl
	bbk_graph --tmpl graph.tmpl > graph.graphviz
graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf
force.csv: graph.csv *.tmpl
	bbk_graph --tmpl force.tmpl > force.csv
force.json: graph.csv *.tmpl
	bbk_graph --tmpl force.json.tmpl > force.json
o: graph.pdf
	open graph.pdf
check:
	bbk_check
clip:
	bbk_clip
clean:
	rm graph.csv
