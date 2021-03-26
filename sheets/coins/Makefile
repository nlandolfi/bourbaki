# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make coins.pdf
	make graph.pdf

coins.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > coins.tex

coins.pdf: ../../*.tex ../../trademark.pdf *.tex coins.tex
	pdflatex --file-line-error -interaction=nonstopmode coins.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: coins.pdf
	open coins.pdf

o:
	make open

defs: coins.tex
	defs coins.tex

spell:
	aspell -c coins.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm coins.*
	make reset
	bbk_sheet -mode mk > Makefile