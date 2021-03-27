# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make conditional_densities.pdf
	make graph.pdf

conditional_densities.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > conditional_densities.tex

conditional_densities.pdf: ../../*.tex ../../trademark.pdf *.tex conditional_densities.tex
	pdflatex --file-line-error -interaction=nonstopmode conditional_densities.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: conditional_densities.pdf
	open conditional_densities.pdf

o:
	make open

defs: conditional_densities.tex
	defs conditional_densities.tex

spell:
	aspell -c conditional_densities.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f conditional_densities.*
	make reset
	bbk_sheet -mode mk > Makefile