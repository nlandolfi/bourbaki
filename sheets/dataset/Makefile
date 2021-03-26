# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make dataset.pdf
	make graph.pdf

dataset.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > dataset.tex

dataset.pdf: ../../*.tex ../../trademark.pdf *.tex dataset.tex
	pdflatex --file-line-error -interaction=nonstopmode dataset.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: dataset.pdf
	open dataset.pdf

o:
	make open

defs: dataset.tex
	defs dataset.tex

spell:
	aspell -c dataset.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm dataset.*
	make reset
	bbk_sheet -mode mk > Makefile