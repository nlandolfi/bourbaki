# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph_cliques.pdf
	make graph.pdf

graph_cliques.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > graph_cliques.tex

graph_cliques.pdf: ../../*.tex ../../trademark.pdf *.tex graph_cliques.tex
	pdflatex --file-line-error -interaction=nonstopmode graph_cliques.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: graph_cliques.pdf
	open graph_cliques.pdf

o:
	make open

defs: graph_cliques.tex
	defs graph_cliques.tex

spell:
	aspell -c graph_cliques.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm graph_cliques.*
	make reset
	bbk_sheet -mode mk > Makefile