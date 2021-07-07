# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix_determinant_of_product.pdf

matrix_determinant_of_product.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_determinant_of_product.tex

matrix_determinant_of_product.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_determinant_of_product.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode matrix_determinant_of_product.tex
	../../latexrun matrix_determinant_of_product.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_determinant_of_product.pdf
	open matrix_determinant_of_product.pdf

o:
	make open

terms: matrix_determinant_of_product.tex
	bbk_sheet -mode ts matrix_determinant_of_product.tex

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
	rm -f matrix_determinant_of_product.*
	make reset
	bbk_sheet -mode mk > Makefile