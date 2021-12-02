# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make diagonally_constant_matrices.pdf

diagonally_constant_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > diagonally_constant_matrices.tex

diagonally_constant_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex diagonally_constant_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode diagonally_constant_matrices.tex
	../../latexrun diagonally_constant_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: diagonally_constant_matrices.pdf
	open diagonally_constant_matrices.pdf

o:
	make open

terms: diagonally_constant_matrices.tex
	bbk_sheet -mode ts diagonally_constant_matrices.tex

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
	rm -f diagonally_constant_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile