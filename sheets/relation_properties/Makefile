# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make relation_properties.pdf
	make graph.pdf

relation_properties.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > relation_properties.tex

relation_properties.pdf: ../../*.tex ../../trademark.pdf *.tex relation_properties.tex
	pdflatex --file-line-error -interaction=nonstopmode relation_properties.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: relation_properties.pdf
	open relation_properties.pdf

o:
	make open

defs: relation_properties.tex
	defs relation_properties.tex

spell:
	aspell -c relation_properties.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm relation_properties.*
	make reset
	bbk_sheet -mode mk > Makefile