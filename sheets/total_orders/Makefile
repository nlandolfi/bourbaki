# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make total_orders.pdf
	make graph.pdf

total_orders.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > total_orders.tex

total_orders.pdf: ../../*.tex ../../trademark.pdf *.tex total_orders.tex
	pdflatex --file-line-error -interaction=nonstopmode total_orders.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: total_orders.pdf
	open total_orders.pdf

o:
	make open

defs: total_orders.tex
	defs total_orders.tex

spell:
	aspell -c total_orders.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm total_orders.*
	make reset
	bbk_sheet -mode mk > Makefile