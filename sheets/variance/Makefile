# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make variance.pdf
	make graph.pdf

variance.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > variance.tex

variance.pdf: ../../*.tex ../../trademark.pdf *.tex variance.tex
	pdflatex --file-line-error -interaction=nonstopmode variance.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: variance.pdf
	open variance.pdf

o:
	make open

defs: variance.tex
	defs variance.tex

spell:
	aspell -c variance.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f variance.*
	make reset
	bbk_sheet -mode mk > Makefile