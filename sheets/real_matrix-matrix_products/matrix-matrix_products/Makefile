# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix-matrix_products.pdf

matrix-matrix_products.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix-matrix_products.tex

matrix-matrix_products.pdf: ../../*.tex ../../trademark.pdf *.tex matrix-matrix_products.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode matrix-matrix_products.tex
	../../latexrun matrix-matrix_products.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix-matrix_products.pdf
	open matrix-matrix_products.pdf

o:
	make open

terms: matrix-matrix_products.tex
	bbk_sheet -mode ts matrix-matrix_products.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f matrix-matrix_products.*
	make reset
	bbk_sheet -mode mk > Makefile