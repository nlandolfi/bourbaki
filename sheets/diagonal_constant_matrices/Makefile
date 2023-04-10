# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make diagonal_constant_matrices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

diagonal_constant_matrices.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > diagonal_constant_matrices.tex

diagonal_constant_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex diagonal_constant_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode diagonal_constant_matrices.tex
	../../latexrun diagonal_constant_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: diagonal_constant_matrices.pdf
	open diagonal_constant_matrices.pdf

o:
	make open

terms: diagonal_constant_matrices.tex
	bbk sheet -mode ts diagonal_constant_matrices.tex

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
	rm -f diagonal_constant_matrices.*
	make reset
	bbk sheet -mode mk > Makefile