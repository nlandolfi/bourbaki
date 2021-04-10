# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make matrix_multiplication.pdf
	make graph.pdf

matrix_multiplication.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_multiplication.tex

matrix_multiplication.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_multiplication.tex
	# pdflatex --file-line-error -interaction=nonstopmode matrix_multiplication.tex
	../../latexrun matrix_multiplication.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_multiplication.pdf
	open matrix_multiplication.pdf

o:
	make open

terms: matrix_multiplication.tex
	bbk_sheet -mode ts matrix_multiplication.tex

spell:
	aspell -c matrix_multiplication.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f matrix_multiplication.*
	make reset
	bbk_sheet -mode mk > Makefile