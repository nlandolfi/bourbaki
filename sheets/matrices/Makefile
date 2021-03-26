# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make matrices.pdf
	make graph.pdf

matrices.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrices.tex

matrices.pdf: ../../*.tex ../../trademark.pdf *.tex matrices.tex
	pdflatex --file-line-error -interaction=nonstopmode matrices.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrices.pdf
	open matrices.pdf

o:
	make open

defs: matrices.tex
	defs matrices.tex

spell:
	aspell -c matrices.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm matrices.*
	make reset
	bbk_sheet -mode mk > Makefile