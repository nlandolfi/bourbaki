# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graphs.pdf
	make graph.pdf

graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > graphs.tex

graphs.pdf: ../../*.tex ../../trademark.pdf *.tex graphs.tex
	pdflatex --file-line-error -interaction=nonstopmode graphs.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: graphs.pdf
	open graphs.pdf

o:
	make open

defs: graphs.tex
	defs graphs.tex

spell:
	aspell -c graphs.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm graphs.*
	make reset
	bbk_sheet -mode mk > Makefile