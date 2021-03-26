# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make bounded_functions.pdf
	make graph.pdf

bounded_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > bounded_functions.tex

bounded_functions.pdf: ../../*.tex ../../trademark.pdf *.tex bounded_functions.tex
	pdflatex --file-line-error -interaction=nonstopmode bounded_functions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: bounded_functions.pdf
	open bounded_functions.pdf

o:
	make open

defs: bounded_functions.tex
	defs bounded_functions.tex

spell:
	aspell -c bounded_functions.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm bounded_functions.*
	make reset
	bbk_sheet -mode mk > Makefile