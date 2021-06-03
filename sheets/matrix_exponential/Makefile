# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make matrix_exponential.pdf

matrix_exponential.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_exponential.tex

matrix_exponential.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_exponential.tex
	# pdflatex --file-line-error -interaction=nonstopmode matrix_exponential.tex
	../../latexrun matrix_exponential.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_exponential.pdf
	open matrix_exponential.pdf

o:
	make open

terms: matrix_exponential.tex
	bbk_sheet -mode ts matrix_exponential.tex

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
	rm -f matrix_exponential.*
	make reset
	bbk_sheet -mode mk > Makefile