# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_matrices_with_orthonormal_columns.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_matrices_with_orthonormal_columns.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_matrices_with_orthonormal_columns.tex

real_matrices_with_orthonormal_columns.pdf: ../../*.tex ../../trademark.pdf *.tex real_matrices_with_orthonormal_columns.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_matrices_with_orthonormal_columns.tex
	../../latexrun real_matrices_with_orthonormal_columns.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_matrices_with_orthonormal_columns.pdf
	open real_matrices_with_orthonormal_columns.pdf

o:
	make open

terms: real_matrices_with_orthonormal_columns.tex
	bbk_sheet -mode ts real_matrices_with_orthonormal_columns.tex

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
	rm -f real_matrices_with_orthonormal_columns.*
	make reset
	bbk_sheet -mode mk > Makefile