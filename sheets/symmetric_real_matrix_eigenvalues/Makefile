# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make symmetric_real_matrix_eigenvalues.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

symmetric_real_matrix_eigenvalues.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > symmetric_real_matrix_eigenvalues.tex

symmetric_real_matrix_eigenvalues.pdf: ../../*.tex ../../trademark.pdf *.tex symmetric_real_matrix_eigenvalues.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode symmetric_real_matrix_eigenvalues.tex
	../../latexrun symmetric_real_matrix_eigenvalues.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: symmetric_real_matrix_eigenvalues.pdf
	open symmetric_real_matrix_eigenvalues.pdf

o:
	make open

terms: symmetric_real_matrix_eigenvalues.tex
	bbk sheet -mode ts symmetric_real_matrix_eigenvalues.tex

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
	rm -f symmetric_real_matrix_eigenvalues.*
	make reset
	bbk sheet -mode mk > Makefile