# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix_space.pdf

matrix_space.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_space.tex

matrix_space.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_space.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode matrix_space.tex
	../../latexrun matrix_space.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_space.pdf
	open matrix_space.pdf

o:
	make open

terms: matrix_space.tex
	bbk_sheet -mode ts matrix_space.tex

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
	rm -f matrix_space.*
	make reset
	bbk_sheet -mode mk > Makefile