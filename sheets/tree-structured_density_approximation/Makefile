# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make tree-structured_density_approximation.pdf
	make graph.pdf

tree-structured_density_approximation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > tree-structured_density_approximation.tex

tree-structured_density_approximation.pdf: ../../*.tex ../../trademark.pdf *.tex tree-structured_density_approximation.tex
	pdflatex --file-line-error -interaction=nonstopmode tree-structured_density_approximation.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tree-structured_density_approximation.pdf
	open tree-structured_density_approximation.pdf

o:
	make open

defs: tree-structured_density_approximation.tex
	defs tree-structured_density_approximation.tex

spell:
	aspell -c tree-structured_density_approximation.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f tree-structured_density_approximation.*
	make reset
	bbk_sheet -mode mk > Makefile