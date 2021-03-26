# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make span.pdf
	make graph.pdf

span.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > span.tex

span.pdf: ../../*.tex ../../trademark.pdf *.tex span.tex
	pdflatex --file-line-error -interaction=nonstopmode span.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: span.pdf
	open span.pdf

o:
	make open

defs: span.tex
	defs span.tex

spell:
	aspell -c span.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm span.*
	make reset
	bbk_sheet -mode mk > Makefile