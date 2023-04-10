# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_matrix_rank.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_matrix_rank.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > real_matrix_rank.tex

real_matrix_rank.pdf: ../../*.tex ../../trademark.pdf *.tex real_matrix_rank.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_matrix_rank.tex
	../../latexrun real_matrix_rank.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_matrix_rank.pdf
	open real_matrix_rank.pdf

o:
	make open

terms: real_matrix_rank.tex
	bbk sheet -mode ts real_matrix_rank.tex

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
	rm -f real_matrix_rank.*
	make reset
	bbk sheet -mode mk > Makefile