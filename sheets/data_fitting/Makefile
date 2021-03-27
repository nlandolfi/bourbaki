# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make data_fitting.pdf
	make graph.pdf

data_fitting.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > data_fitting.tex

data_fitting.pdf: ../../*.tex ../../trademark.pdf *.tex data_fitting.tex
	pdflatex --file-line-error -interaction=nonstopmode data_fitting.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: data_fitting.pdf
	open data_fitting.pdf

o:
	make open

defs: data_fitting.tex
	defs data_fitting.tex

spell:
	aspell -c data_fitting.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f data_fitting.*
	make reset
	bbk_sheet -mode mk > Makefile