# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make trees.pdf
	make graph.pdf

trees.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > trees.tex

trees.pdf: ../../*.tex ../../trademark.pdf *.tex trees.tex
	pdflatex --file-line-error -interaction=nonstopmode trees.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: trees.pdf
	open trees.pdf

o:
	make open

defs: trees.tex
	defs trees.tex

spell:
	aspell -c trees.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm trees.*
	make reset
	bbk_sheet -mode mk > Makefile