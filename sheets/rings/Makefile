# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make rings.pdf
	make graph.pdf

rings.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > rings.tex

rings.pdf: ../../*.tex ../../trademark.pdf *.tex rings.tex
	pdflatex --file-line-error -interaction=nonstopmode rings.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: rings.pdf
	open rings.pdf

o:
	make open

defs: rings.tex
	defs rings.tex

spell:
	aspell -c rings.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f rings.*
	make reset
	bbk_sheet -mode mk > Makefile