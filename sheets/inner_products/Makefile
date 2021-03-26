# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make inner_products.pdf
	make graph.pdf

inner_products.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > inner_products.tex

inner_products.pdf: ../../*.tex ../../trademark.pdf *.tex inner_products.tex
	pdflatex --file-line-error -interaction=nonstopmode inner_products.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inner_products.pdf
	open inner_products.pdf

o:
	make open

defs: inner_products.tex
	defs inner_products.tex

spell:
	aspell -c inner_products.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm inner_products.*
	make reset
	bbk_sheet -mode mk > Makefile