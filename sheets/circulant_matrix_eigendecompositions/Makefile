# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make circulant_matrix_eigendecompositions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

circulant_matrix_eigendecompositions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > circulant_matrix_eigendecompositions.tex

circulant_matrix_eigendecompositions.pdf: ../../*.tex ../../trademark.pdf *.tex circulant_matrix_eigendecompositions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode circulant_matrix_eigendecompositions.tex
	../../latexrun circulant_matrix_eigendecompositions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: circulant_matrix_eigendecompositions.pdf
	open circulant_matrix_eigendecompositions.pdf

o:
	make open

terms: circulant_matrix_eigendecompositions.tex
	bbk sheet -mode ts circulant_matrix_eigendecompositions.tex

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
	bbk sheet -mode mk > Makefile