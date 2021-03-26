# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make rooted_trees.pdf
	make graph.pdf

rooted_trees.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > rooted_trees.tex

rooted_trees.pdf: ../../*.tex ../../trademark.pdf *.tex rooted_trees.tex
	pdflatex --file-line-error -interaction=nonstopmode rooted_trees.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rooted_trees.pdf
	open rooted_trees.pdf

o:
	make open

defs: rooted_trees.tex
	defs rooted_trees.tex

spell:
	aspell -c rooted_trees.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm rooted_trees.*
	make reset
	bbk_sheet -mode mk > Makefile