# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make data_matrix.pdf
	make graph.pdf

data_matrix.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > data_matrix.tex

data_matrix.pdf: ../../*.tex ../../trademark.pdf *.tex data_matrix.tex
	pdflatex --file-line-error -interaction=nonstopmode data_matrix.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: data_matrix.pdf
	open data_matrix.pdf

o:
	make open

defs: data_matrix.tex
	defs data_matrix.tex

spell:
	aspell -c data_matrix.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm data_matrix.*
	make reset
	bbk_sheet -mode mk > Makefile