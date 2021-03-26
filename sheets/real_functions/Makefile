# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make real_functions.pdf
	make graph.pdf

real_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_functions.tex

real_functions.pdf: ../../*.tex ../../trademark.pdf *.tex real_functions.tex
	pdflatex --file-line-error -interaction=nonstopmode real_functions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_functions.pdf
	open real_functions.pdf

o:
	make open

defs: real_functions.tex
	defs real_functions.tex

spell:
	aspell -c real_functions.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm real_functions.*
	make reset
	bbk_sheet -mode mk > Makefile