# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make block_diagonal_matrices.pdf

block_diagonal_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > block_diagonal_matrices.tex

block_diagonal_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex block_diagonal_matrices.tex
	# pdflatex --file-line-error -interaction=nonstopmode block_diagonal_matrices.tex
	../../latexrun block_diagonal_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: block_diagonal_matrices.pdf
	open block_diagonal_matrices.pdf

o:
	make open

terms: block_diagonal_matrices.tex
	bbk_sheet -mode ts block_diagonal_matrices.tex

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
	rm -f block_diagonal_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile