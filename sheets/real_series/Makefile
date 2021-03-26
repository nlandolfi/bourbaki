# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make real_series.pdf
	make graph.pdf

real_series.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_series.tex

real_series.pdf: ../../*.tex ../../trademark.pdf *.tex real_series.tex
	pdflatex --file-line-error -interaction=nonstopmode real_series.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_series.pdf
	open real_series.pdf

o:
	make open

defs: real_series.tex
	defs real_series.tex

spell:
	aspell -c real_series.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm real_series.*
	make reset
	bbk_sheet -mode mk > Makefile