# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make tree-structured_distributions.pdf
	make graph.pdf

tree-structured_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > tree-structured_distributions.tex

tree-structured_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex tree-structured_distributions.tex
	pdflatex --file-line-error -interaction=nonstopmode tree-structured_distributions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tree-structured_distributions.pdf
	open tree-structured_distributions.pdf

o:
	make open

defs: tree-structured_distributions.tex
	defs tree-structured_distributions.tex

spell:
	aspell -c tree-structured_distributions.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm tree-structured_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile