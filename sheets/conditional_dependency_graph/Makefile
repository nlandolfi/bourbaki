# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make conditional_dependency_graph.pdf
	make graph.pdf

conditional_dependency_graph.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > conditional_dependency_graph.tex

conditional_dependency_graph.pdf: ../../*.tex ../../trademark.pdf *.tex conditional_dependency_graph.tex
	pdflatex --file-line-error -interaction=nonstopmode conditional_dependency_graph.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: conditional_dependency_graph.pdf
	open conditional_dependency_graph.pdf

o:
	make open

defs: conditional_dependency_graph.tex
	defs conditional_dependency_graph.tex

spell:
	aspell -c conditional_dependency_graph.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm conditional_dependency_graph.*
	make reset
	bbk_sheet -mode mk > Makefile