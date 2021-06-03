# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix_transpose.pdf

matrix_transpose.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_transpose.tex

matrix_transpose.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_transpose.tex
	# pdflatex --file-line-error -interaction=nonstopmode matrix_transpose.tex
	../../latexrun matrix_transpose.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_transpose.pdf
	open matrix_transpose.pdf

o:
	make open

terms: matrix_transpose.tex
	bbk_sheet -mode ts matrix_transpose.tex

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
	rm -f matrix_transpose.*
	make reset
	bbk_sheet -mode mk > Makefile