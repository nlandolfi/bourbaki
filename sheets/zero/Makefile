# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make zero.pdf
	make graph.pdf

zero.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > zero.tex

zero.pdf: ../../*.tex ../../trademark.pdf *.tex zero.tex
	pdflatex --file-line-error -interaction=nonstopmode zero.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: zero.pdf
	open zero.pdf

o:
	make open

defs: zero.tex
	defs zero.tex

spell:
	aspell -c zero.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm zero.*
	make reset
	bbk_sheet -mode mk > Makefile