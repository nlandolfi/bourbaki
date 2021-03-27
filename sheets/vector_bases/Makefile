# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make vector_bases.pdf
	make graph.pdf

vector_bases.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > vector_bases.tex

vector_bases.pdf: ../../*.tex ../../trademark.pdf *.tex vector_bases.tex
	pdflatex --file-line-error -interaction=nonstopmode vector_bases.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: vector_bases.pdf
	open vector_bases.pdf

o:
	make open

defs: vector_bases.tex
	defs vector_bases.tex

spell:
	aspell -c vector_bases.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f vector_bases.*
	make reset
	bbk_sheet -mode mk > Makefile