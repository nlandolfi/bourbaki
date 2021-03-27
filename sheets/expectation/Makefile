# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make expectation.pdf
	make graph.pdf

expectation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > expectation.tex

expectation.pdf: ../../*.tex ../../trademark.pdf *.tex expectation.tex
	pdflatex --file-line-error -interaction=nonstopmode expectation.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: expectation.pdf
	open expectation.pdf

o:
	make open

defs: expectation.tex
	defs expectation.tex

spell:
	aspell -c expectation.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f expectation.*
	make reset
	bbk_sheet -mode mk > Makefile