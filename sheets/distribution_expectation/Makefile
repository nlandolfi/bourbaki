# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make distribution_expectation.pdf
	make graph.pdf

distribution_expectation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > distribution_expectation.tex

distribution_expectation.pdf: ../../*.tex ../../trademark.pdf *.tex distribution_expectation.tex
	pdflatex --file-line-error -interaction=nonstopmode distribution_expectation.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distribution_expectation.pdf
	open distribution_expectation.pdf

o:
	make open

defs: distribution_expectation.tex
	defs distribution_expectation.tex

spell:
	aspell -c distribution_expectation.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm distribution_expectation.*
	make reset
	bbk_sheet -mode mk > Makefile