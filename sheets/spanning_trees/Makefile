# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make spanning_trees.pdf
	make graph.pdf

spanning_trees.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > spanning_trees.tex

spanning_trees.pdf: ../../*.tex ../../trademark.pdf *.tex spanning_trees.tex
	# pdflatex --file-line-error -interaction=nonstopmode spanning_trees.tex
	../../latexrun spanning_trees.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: spanning_trees.pdf
	open spanning_trees.pdf

o:
	make open

terms: spanning_trees.tex
	bbk_sheet -mode ts spanning_trees.tex

spell:
	aspell -c spanning_trees.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f spanning_trees.*
	make reset
	bbk_sheet -mode mk > Makefile