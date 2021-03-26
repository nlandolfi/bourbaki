# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make set_extension.pdf
	make graph.pdf

set_extension.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_extension.tex

set_extension.pdf: ../../*.tex ../../trademark.pdf *.tex set_extension.tex
	pdflatex --file-line-error -interaction=nonstopmode set_extension.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_extension.pdf
	open set_extension.pdf

o:
	make open

defs: set_extension.tex
	defs set_extension.tex

spell:
	aspell -c set_extension.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm set_extension.*
	make reset
	bbk_sheet -mode mk > Makefile