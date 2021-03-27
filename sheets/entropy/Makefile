# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make entropy.pdf
	make graph.pdf

entropy.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > entropy.tex

entropy.pdf: ../../*.tex ../../trademark.pdf *.tex entropy.tex
	pdflatex --file-line-error -interaction=nonstopmode entropy.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: entropy.pdf
	open entropy.pdf

o:
	make open

defs: entropy.tex
	defs entropy.tex

spell:
	aspell -c entropy.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f entropy.*
	make reset
	bbk_sheet -mode mk > Makefile