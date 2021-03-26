# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make cross_entropy.pdf
	make graph.pdf

cross_entropy.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cross_entropy.tex

cross_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex cross_entropy.tex
	pdflatex --file-line-error -interaction=nonstopmode cross_entropy.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cross_entropy.pdf
	open cross_entropy.pdf

o:
	make open

defs: cross_entropy.tex
	defs cross_entropy.tex

spell:
	aspell -c cross_entropy.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm cross_entropy.*
	make reset
	bbk_sheet -mode mk > Makefile