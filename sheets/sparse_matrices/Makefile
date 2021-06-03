# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make sparse_matrices.pdf

sparse_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sparse_matrices.tex

sparse_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex sparse_matrices.tex
	# pdflatex --file-line-error -interaction=nonstopmode sparse_matrices.tex
	../../latexrun sparse_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sparse_matrices.pdf
	open sparse_matrices.pdf

o:
	make open

terms: sparse_matrices.tex
	bbk_sheet -mode ts sparse_matrices.tex

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
	rm -f sparse_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile