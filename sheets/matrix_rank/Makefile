# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix_rank.pdf

matrix_rank.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_rank.tex

matrix_rank.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_rank.tex
	# pdflatex --file-line-error -interaction=nonstopmode matrix_rank.tex
	../../latexrun matrix_rank.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_rank.pdf
	open matrix_rank.pdf

o:
	make open

terms: matrix_rank.tex
	bbk_sheet -mode ts matrix_rank.tex

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
	rm -f matrix_rank.*
	make reset
	bbk_sheet -mode mk > Makefile