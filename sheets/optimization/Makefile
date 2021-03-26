# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make optimization.pdf
	make graph.pdf

optimization.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > optimization.tex

optimization.pdf: ../../*.tex ../../trademark.pdf *.tex optimization.tex
	pdflatex --file-line-error -interaction=nonstopmode optimization.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: optimization.pdf
	open optimization.pdf

o:
	make open

defs: optimization.tex
	defs optimization.tex

spell:
	aspell -c optimization.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm optimization.*
	make reset
	bbk_sheet -mode mk > Makefile