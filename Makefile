all: */*.tex
	cd all/ && make open

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

graph: graph.pdf
	make graph.pdf
	open graph.pdf

open:
	make graph

o:
	make open

clean:
	rm */*.aux
	rm */*.out
	rm */*.log


