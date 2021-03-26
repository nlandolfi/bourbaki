# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make triangle_equality.pdf
	make graph.pdf

triangle_equality.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > triangle_equality.tex

triangle_equality.pdf: ../../*.tex ../../trademark.pdf *.tex triangle_equality.tex
	pdflatex --file-line-error -interaction=nonstopmode triangle_equality.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: triangle_equality.pdf
	open triangle_equality.pdf

o:
	make open

defs: triangle_equality.tex
	defs triangle_equality.tex

spell:
	aspell -c triangle_equality.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm triangle_equality.*
	make reset
	bbk_sheet -mode mk > Makefile