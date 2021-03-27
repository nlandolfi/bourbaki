# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make partial_orders.pdf
	make graph.pdf

partial_orders.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > partial_orders.tex

partial_orders.pdf: ../../*.tex ../../trademark.pdf *.tex partial_orders.tex
	pdflatex --file-line-error -interaction=nonstopmode partial_orders.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: partial_orders.pdf
	open partial_orders.pdf

o:
	make open

defs: partial_orders.tex
	defs partial_orders.tex

spell:
	aspell -c partial_orders.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f partial_orders.*
	make reset
	bbk_sheet -mode mk > Makefile