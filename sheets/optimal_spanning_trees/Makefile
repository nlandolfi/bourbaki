# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make optimal_spanning_trees.pdf
	make graph.pdf

optimal_spanning_trees.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > optimal_spanning_trees.tex

optimal_spanning_trees.pdf: ../../*.tex ../../trademark.pdf *.tex optimal_spanning_trees.tex
	# pdflatex --file-line-error -interaction=nonstopmode optimal_spanning_trees.tex
	../../latexrun optimal_spanning_trees.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: optimal_spanning_trees.pdf
	open optimal_spanning_trees.pdf

o:
	make open

terms: optimal_spanning_trees.tex
	bbk_sheet -mode ts optimal_spanning_trees.tex

spell:
	aspell -c optimal_spanning_trees.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f optimal_spanning_trees.*
	make reset
	bbk_sheet -mode mk > Makefile