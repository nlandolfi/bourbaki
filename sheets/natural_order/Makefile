# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make natural_order.pdf
	make graph.pdf

natural_order.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_order.tex

natural_order.pdf: ../../*.tex ../../trademark.pdf *.tex natural_order.tex
	pdflatex --file-line-error -interaction=nonstopmode natural_order.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_order.pdf
	open natural_order.pdf

o:
	make open

defs: natural_order.tex
	defs natural_order.tex

spell:
	aspell -c natural_order.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f natural_order.*
	make reset
	bbk_sheet -mode mk > Makefile