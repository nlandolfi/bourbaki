# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make complete_inner_product_decomposition.pdf
	make graph.pdf

complete_inner_product_decomposition.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complete_inner_product_decomposition.tex

complete_inner_product_decomposition.pdf: ../../*.tex ../../trademark.pdf *.tex complete_inner_product_decomposition.tex
	pdflatex --file-line-error -interaction=nonstopmode complete_inner_product_decomposition.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complete_inner_product_decomposition.pdf
	open complete_inner_product_decomposition.pdf

o:
	make open

defs: complete_inner_product_decomposition.tex
	defs complete_inner_product_decomposition.tex

spell:
	aspell -c complete_inner_product_decomposition.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm complete_inner_product_decomposition.*
	make reset
	bbk_sheet -mode mk > Makefile