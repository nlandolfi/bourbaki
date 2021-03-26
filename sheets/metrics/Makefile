# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make metrics.pdf
	make graph.pdf

metrics.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metrics.tex

metrics.pdf: ../../*.tex ../../trademark.pdf *.tex metrics.tex
	pdflatex --file-line-error -interaction=nonstopmode metrics.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metrics.pdf
	open metrics.pdf

o:
	make open

defs: metrics.tex
	defs metrics.tex

spell:
	aspell -c metrics.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm metrics.*
	make reset
	bbk_sheet -mode mk > Makefile