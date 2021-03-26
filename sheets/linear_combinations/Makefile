# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make linear_combinations.pdf
	make graph.pdf

linear_combinations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > linear_combinations.tex

linear_combinations.pdf: ../../*.tex ../../trademark.pdf *.tex linear_combinations.tex
	pdflatex --file-line-error -interaction=nonstopmode linear_combinations.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_combinations.pdf
	open linear_combinations.pdf

o:
	make open

defs: linear_combinations.tex
	defs linear_combinations.tex

spell:
	aspell -c linear_combinations.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm linear_combinations.*
	make reset
	bbk_sheet -mode mk > Makefile