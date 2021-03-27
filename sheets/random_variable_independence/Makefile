# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make random_variable_independence.pdf
	make graph.pdf

random_variable_independence.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > random_variable_independence.tex

random_variable_independence.pdf: ../../*.tex ../../trademark.pdf *.tex random_variable_independence.tex
	pdflatex --file-line-error -interaction=nonstopmode random_variable_independence.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: random_variable_independence.pdf
	open random_variable_independence.pdf

o:
	make open

defs: random_variable_independence.tex
	defs random_variable_independence.tex

spell:
	aspell -c random_variable_independence.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f random_variable_independence.*
	make reset
	bbk_sheet -mode mk > Makefile