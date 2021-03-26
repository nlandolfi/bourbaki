# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make vector_space_dimension.pdf
	make graph.pdf

vector_space_dimension.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > vector_space_dimension.tex

vector_space_dimension.pdf: ../../*.tex ../../trademark.pdf *.tex vector_space_dimension.tex
	pdflatex --file-line-error -interaction=nonstopmode vector_space_dimension.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: vector_space_dimension.pdf
	open vector_space_dimension.pdf

o:
	make open

defs: vector_space_dimension.tex
	defs vector_space_dimension.tex

spell:
	aspell -c vector_space_dimension.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm vector_space_dimension.*
	make reset
	bbk_sheet -mode mk > Makefile