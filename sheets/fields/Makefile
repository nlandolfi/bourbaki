# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make fields.pdf
	make graph.pdf

fields.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > fields.tex

fields.pdf: ../../*.tex ../../trademark.pdf *.tex fields.tex
	pdflatex --file-line-error -interaction=nonstopmode fields.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: fields.pdf
	open fields.pdf

o:
	make open

defs: fields.tex
	defs fields.tex

spell:
	aspell -c fields.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm fields.*
	make reset
	bbk_sheet -mode mk > Makefile