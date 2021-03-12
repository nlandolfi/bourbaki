# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make objects.pdf
	make graph.pdf

objects.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > objects.tex

objects.pdf: ../../*.tex ../../trademark.pdf *.tex objects.tex
	pdflatex objects.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: objects.pdf
	open objects.pdf

o:
	make open

defs: objects.tex
	defs objects.tex

spell:
	aspell -c objects.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm objects.*
	make reset
	bbk_sheet -mode mk > Makefile