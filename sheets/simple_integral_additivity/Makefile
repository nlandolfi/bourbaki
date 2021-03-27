# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make simple_integral_additivity.pdf
	make graph.pdf

simple_integral_additivity.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > simple_integral_additivity.tex

simple_integral_additivity.pdf: ../../*.tex ../../trademark.pdf *.tex simple_integral_additivity.tex
	pdflatex --file-line-error -interaction=nonstopmode simple_integral_additivity.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: simple_integral_additivity.pdf
	open simple_integral_additivity.pdf

o:
	make open

defs: simple_integral_additivity.tex
	defs simple_integral_additivity.tex

spell:
	aspell -c simple_integral_additivity.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f simple_integral_additivity.*
	make reset
	bbk_sheet -mode mk > Makefile