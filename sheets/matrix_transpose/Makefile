# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix_transpose.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

matrix_transpose.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > matrix_transpose.tex

matrix_transpose.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_transpose.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode matrix_transpose.tex
	../../latexrun matrix_transpose.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_transpose.pdf
	open matrix_transpose.pdf

o:
	make open

terms: matrix_transpose.tex
	bbk sheet -mode ts matrix_transpose.tex

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
	bbk sheet -mode mk > Makefile