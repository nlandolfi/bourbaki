# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix_squares.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

matrix_squares.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > matrix_squares.tex

matrix_squares.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_squares.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode matrix_squares.tex
	../../latexrun matrix_squares.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_squares.pdf
	open matrix_squares.pdf

o:
	make open

terms: matrix_squares.tex
	bbk_sheet -mode ts matrix_squares.tex

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
	rm -f matrix_squares.*
	make reset
	bbk_sheet -mode mk > Makefile