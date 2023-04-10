# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make normal_matrices.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

normal_matrices.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > normal_matrices.tex

normal_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex normal_matrices.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode normal_matrices.tex
	../../latexrun normal_matrices.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: normal_matrices.pdf
	open normal_matrices.pdf

o:
	make open

terms: normal_matrices.tex
	bbk sheet -mode ts normal_matrices.tex

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
	rm -f normal_matrices.*
	make reset
	bbk sheet -mode mk > Makefile