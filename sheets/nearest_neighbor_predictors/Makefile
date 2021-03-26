# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make nearest_neighbor_predictors.pdf
	make graph.pdf

nearest_neighbor_predictors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > nearest_neighbor_predictors.tex

nearest_neighbor_predictors.pdf: ../../*.tex ../../trademark.pdf *.tex nearest_neighbor_predictors.tex
	pdflatex --file-line-error -interaction=nonstopmode nearest_neighbor_predictors.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: nearest_neighbor_predictors.pdf
	open nearest_neighbor_predictors.pdf

o:
	make open

defs: nearest_neighbor_predictors.tex
	defs nearest_neighbor_predictors.tex

spell:
	aspell -c nearest_neighbor_predictors.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm nearest_neighbor_predictors.*
	make reset
	bbk_sheet -mode mk > Makefile