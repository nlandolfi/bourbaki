# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make partial_derivatives.pdf
	make graph.pdf

partial_derivatives.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > partial_derivatives.tex

partial_derivatives.pdf: ../../*.tex ../../trademark.pdf *.tex partial_derivatives.tex
	pdflatex --file-line-error -interaction=nonstopmode partial_derivatives.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: partial_derivatives.pdf
	open partial_derivatives.pdf

o:
	make open

defs: partial_derivatives.tex
	defs partial_derivatives.tex

spell:
	aspell -c partial_derivatives.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm partial_derivatives.*
	make reset
	bbk_sheet -mode mk > Makefile