# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make matrix_eigenvectors.pdf
	make graph.pdf

matrix_eigenvectors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_eigenvectors.tex

matrix_eigenvectors.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_eigenvectors.tex
	# pdflatex --file-line-error -interaction=nonstopmode matrix_eigenvectors.tex
	../../latexrun matrix_eigenvectors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_eigenvectors.pdf
	open matrix_eigenvectors.pdf

o:
	make open

terms: matrix_eigenvectors.tex
	bbk_sheet -mode ts matrix_eigenvectors.tex

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
	rm -f matrix_eigenvectors.*
	make reset
	bbk_sheet -mode mk > Makefile