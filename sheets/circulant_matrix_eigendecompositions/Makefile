# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make circulant_matrix_eigendecompositions.pdf

circulant_matrix_eigendecompositions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > circulant_matrix_eigendecompositions.tex

circulant_matrix_eigendecompositions.pdf: ../../*.tex ../../trademark.pdf *.tex circulant_matrix_eigendecompositions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode circulant_matrix_eigendecompositions.tex
	../../latexrun circulant_matrix_eigendecompositions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: circulant_matrix_eigendecompositions.pdf
	open circulant_matrix_eigendecompositions.pdf

o:
	make open

terms: circulant_matrix_eigendecompositions.tex
	bbk_sheet -mode ts circulant_matrix_eigendecompositions.tex

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
	rm -f circulant_matrix_eigendecompositions.*
	make reset
	bbk_sheet -mode mk > Makefile