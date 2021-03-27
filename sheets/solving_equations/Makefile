# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make solving_equations.pdf
	make graph.pdf

solving_equations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > solving_equations.tex

solving_equations.pdf: ../../*.tex ../../trademark.pdf *.tex solving_equations.tex
	pdflatex --file-line-error -interaction=nonstopmode solving_equations.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: solving_equations.pdf
	open solving_equations.pdf

o:
	make open

defs: solving_equations.tex
	defs solving_equations.tex

spell:
	aspell -c solving_equations.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f solving_equations.*
	make reset
	bbk_sheet -mode mk > Makefile