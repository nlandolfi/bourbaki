# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make identity_matrices.pdf
	make graph.pdf

identity_matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > identity_matrices.tex

identity_matrices.pdf: ../../*.tex ../../trademark.pdf *.tex identity_matrices.tex
	pdflatex --file-line-error -interaction=nonstopmode identity_matrices.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: identity_matrices.pdf
	open identity_matrices.pdf

o:
	make open

defs: identity_matrices.tex
	defs identity_matrices.tex

spell:
	aspell -c identity_matrices.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm identity_matrices.*
	make reset
	bbk_sheet -mode mk > Makefile