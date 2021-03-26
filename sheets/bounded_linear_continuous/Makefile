# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make bounded_linear_continuous.pdf
	make graph.pdf

bounded_linear_continuous.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > bounded_linear_continuous.tex

bounded_linear_continuous.pdf: ../../*.tex ../../trademark.pdf *.tex bounded_linear_continuous.tex
	pdflatex --file-line-error -interaction=nonstopmode bounded_linear_continuous.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: bounded_linear_continuous.pdf
	open bounded_linear_continuous.pdf

o:
	make open

defs: bounded_linear_continuous.tex
	defs bounded_linear_continuous.tex

spell:
	aspell -c bounded_linear_continuous.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm bounded_linear_continuous.*
	make reset
	bbk_sheet -mode mk > Makefile