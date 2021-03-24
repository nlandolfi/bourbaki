# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make maximum_likelihood.pdf
	make graph.pdf

maximum_likelihood.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > maximum_likelihood.tex

maximum_likelihood.pdf: ../../*.tex ../../trademark.pdf *.tex maximum_likelihood.tex
	pdflatex maximum_likelihood.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: maximum_likelihood.pdf
	open maximum_likelihood.pdf

o:
	make open

defs: maximum_likelihood.tex
	defs maximum_likelihood.tex

spell:
	aspell -c maximum_likelihood.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm maximum_likelihood.*
	make reset
	bbk_sheet -mode mk > Makefile