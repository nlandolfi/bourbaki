# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make matrix_determinant_of_inverse.pdf
	make graph.pdf

matrix_determinant_of_inverse.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_determinant_of_inverse.tex

matrix_determinant_of_inverse.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_determinant_of_inverse.tex
	# pdflatex --file-line-error -interaction=nonstopmode matrix_determinant_of_inverse.tex
	../../latexrun matrix_determinant_of_inverse.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_determinant_of_inverse.pdf
	open matrix_determinant_of_inverse.pdf

o:
	make open

terms: matrix_determinant_of_inverse.tex
	bbk_sheet -mode ts matrix_determinant_of_inverse.tex

spell:
	aspell -c matrix_determinant_of_inverse.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f matrix_determinant_of_inverse.*
	make reset
	bbk_sheet -mode mk > Makefile