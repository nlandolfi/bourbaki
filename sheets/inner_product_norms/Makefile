# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make inner_product_norms.pdf
	make graph.pdf

inner_product_norms.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > inner_product_norms.tex

inner_product_norms.pdf: ../../*.tex ../../trademark.pdf *.tex inner_product_norms.tex
	pdflatex --file-line-error -interaction=nonstopmode inner_product_norms.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inner_product_norms.pdf
	open inner_product_norms.pdf

o:
	make open

defs: inner_product_norms.tex
	defs inner_product_norms.tex

spell:
	aspell -c inner_product_norms.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm inner_product_norms.*
	make reset
	bbk_sheet -mode mk > Makefile