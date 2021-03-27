# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make tree-structured_densities.pdf
	make graph.pdf

tree-structured_densities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > tree-structured_densities.tex

tree-structured_densities.pdf: ../../*.tex ../../trademark.pdf *.tex tree-structured_densities.tex
	pdflatex --file-line-error -interaction=nonstopmode tree-structured_densities.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tree-structured_densities.pdf
	open tree-structured_densities.pdf

o:
	make open

defs: tree-structured_densities.tex
	defs tree-structured_densities.tex

spell:
	aspell -c tree-structured_densities.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f tree-structured_densities.*
	make reset
	bbk_sheet -mode mk > Makefile