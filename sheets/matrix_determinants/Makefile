# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make matrix_determinants.pdf
	make graph.pdf

matrix_determinants.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_determinants.tex

matrix_determinants.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_determinants.tex
	pdflatex --file-line-error -interaction=nonstopmode matrix_determinants.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_determinants.pdf
	open matrix_determinants.pdf

o:
	make open

defs: matrix_determinants.tex
	defs matrix_determinants.tex

spell:
	aspell -c matrix_determinants.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm matrix_determinants.*
	make reset
	bbk_sheet -mode mk > Makefile