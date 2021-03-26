# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make complex_inner_products.pdf
	make graph.pdf

complex_inner_products.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_inner_products.tex

complex_inner_products.pdf: ../../*.tex ../../trademark.pdf *.tex complex_inner_products.tex
	pdflatex --file-line-error -interaction=nonstopmode complex_inner_products.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_inner_products.pdf
	open complex_inner_products.pdf

o:
	make open

defs: complex_inner_products.tex
	defs complex_inner_products.tex

spell:
	aspell -c complex_inner_products.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm complex_inner_products.*
	make reset
	bbk_sheet -mode mk > Makefile