# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make real_limiting_bounds.pdf
	make graph.pdf

real_limiting_bounds.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_limiting_bounds.tex

real_limiting_bounds.pdf: ../../*.tex ../../trademark.pdf *.tex real_limiting_bounds.tex
	pdflatex --file-line-error -interaction=nonstopmode real_limiting_bounds.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_limiting_bounds.pdf
	open real_limiting_bounds.pdf

o:
	make open

defs: real_limiting_bounds.tex
	defs real_limiting_bounds.tex

spell:
	aspell -c real_limiting_bounds.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f real_limiting_bounds.*
	make reset
	bbk_sheet -mode mk > Makefile