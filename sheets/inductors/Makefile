# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make inductors.pdf
	make graph.pdf

inductors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > inductors.tex

inductors.pdf: ../../*.tex ../../trademark.pdf *.tex inductors.tex
	pdflatex --file-line-error -interaction=nonstopmode inductors.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inductors.pdf
	open inductors.pdf

o:
	make open

defs: inductors.tex
	defs inductors.tex

spell:
	aspell -c inductors.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm inductors.*
	make reset
	bbk_sheet -mode mk > Makefile