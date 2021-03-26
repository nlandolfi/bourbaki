# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make relations.pdf
	make graph.pdf

relations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > relations.tex

relations.pdf: ../../*.tex ../../trademark.pdf *.tex relations.tex
	pdflatex --file-line-error -interaction=nonstopmode relations.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: relations.pdf
	open relations.pdf

o:
	make open

defs: relations.tex
	defs relations.tex

spell:
	aspell -c relations.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm relations.*
	make reset
	bbk_sheet -mode mk > Makefile