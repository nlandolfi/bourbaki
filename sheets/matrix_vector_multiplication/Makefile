# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make matrix_vector_multiplication.pdf
	make graph.pdf

matrix_vector_multiplication.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_vector_multiplication.tex

matrix_vector_multiplication.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_vector_multiplication.tex
	# pdflatex --file-line-error -interaction=nonstopmode matrix_vector_multiplication.tex
	../../latexrun matrix_vector_multiplication.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_vector_multiplication.pdf
	open matrix_vector_multiplication.pdf

o:
	make open

terms: matrix_vector_multiplication.tex
	bbk_sheet -mode ts matrix_vector_multiplication.tex

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
	rm -f matrix_vector_multiplication.*
	make reset
	bbk_sheet -mode mk > Makefile