# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make expectation_maximization.pdf
	make graph.pdf

expectation_maximization.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > expectation_maximization.tex

expectation_maximization.pdf: ../../*.tex ../../trademark.pdf *.tex expectation_maximization.tex
	pdflatex expectation_maximization.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: expectation_maximization.pdf
	open expectation_maximization.pdf

o:
	make open

defs: expectation_maximization.tex
	defs expectation_maximization.tex

spell:
	aspell -c expectation_maximization.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm expectation_maximization.*
	make reset
	bbk_sheet -mode mk > Makefile