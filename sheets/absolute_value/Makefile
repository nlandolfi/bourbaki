# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make absolute_value.pdf
	make graph.pdf

absolute_value.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > absolute_value.tex

absolute_value.pdf: ../../*.tex ../../trademark.pdf *.tex absolute_value.tex
	pdflatex --file-line-error -interaction=nonstopmode absolute_value.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: absolute_value.pdf
	open absolute_value.pdf

o:
	make open

defs: absolute_value.tex
	defs absolute_value.tex

spell:
	aspell -c absolute_value.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm absolute_value.*
	make reset
	bbk_sheet -mode mk > Makefile