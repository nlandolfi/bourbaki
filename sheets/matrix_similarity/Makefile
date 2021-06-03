# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix_similarity.pdf

matrix_similarity.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_similarity.tex

matrix_similarity.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_similarity.tex
	# pdflatex --file-line-error -interaction=nonstopmode matrix_similarity.tex
	../../latexrun matrix_similarity.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_similarity.pdf
	open matrix_similarity.pdf

o:
	make open

terms: matrix_similarity.tex
	bbk_sheet -mode ts matrix_similarity.tex

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
	rm -f matrix_similarity.*
	make reset
	bbk_sheet -mode mk > Makefile