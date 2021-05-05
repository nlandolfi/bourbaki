# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make matrix_vector_products.pdf
	make graph.pdf

matrix_vector_products.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_vector_products.tex

matrix_vector_products.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_vector_products.tex
	# pdflatex --file-line-error -interaction=nonstopmode matrix_vector_products.tex
	../../latexrun matrix_vector_products.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_vector_products.pdf
	open matrix_vector_products.pdf

o:
	make open

terms: matrix_vector_products.tex
	bbk_sheet -mode ts matrix_vector_products.tex

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
	rm -f matrix_vector_products.*
	make reset
	bbk_sheet -mode mk > Makefile