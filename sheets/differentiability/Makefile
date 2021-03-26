# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make differentiability.pdf
	make graph.pdf

differentiability.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > differentiability.tex

differentiability.pdf: ../../*.tex ../../trademark.pdf *.tex differentiability.tex
	pdflatex --file-line-error -interaction=nonstopmode differentiability.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: differentiability.pdf
	open differentiability.pdf

o:
	make open

defs: differentiability.tex
	defs differentiability.tex

spell:
	aspell -c differentiability.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm differentiability.*
	make reset
	bbk_sheet -mode mk > Makefile