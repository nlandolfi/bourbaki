# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make prefix_trees.pdf

prefix_trees.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > prefix_trees.tex

prefix_trees.pdf: ../../*.tex ../../trademark.pdf *.tex prefix_trees.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode prefix_trees.tex
	../../latexrun prefix_trees.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: prefix_trees.pdf
	open prefix_trees.pdf

o:
	make open

terms: prefix_trees.tex
	bbk_sheet -mode ts prefix_trees.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f prefix_trees.*
	make reset
	bbk_sheet -mode mk > Makefile