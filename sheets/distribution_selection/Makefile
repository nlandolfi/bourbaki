# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make distribution_selection.pdf
	make graph.pdf

distribution_selection.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > distribution_selection.tex

distribution_selection.pdf: ../../*.tex ../../trademark.pdf *.tex distribution_selection.tex
	pdflatex --file-line-error -interaction=nonstopmode distribution_selection.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distribution_selection.pdf
	open distribution_selection.pdf

o:
	make open

defs: distribution_selection.tex
	defs distribution_selection.tex

spell:
	aspell -c distribution_selection.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm distribution_selection.*
	make reset
	bbk_sheet -mode mk > Makefile