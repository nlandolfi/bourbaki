# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make inverse_elements.pdf
	make graph.pdf

inverse_elements.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > inverse_elements.tex

inverse_elements.pdf: ../../*.tex ../../trademark.pdf *.tex inverse_elements.tex
	pdflatex --file-line-error -interaction=nonstopmode inverse_elements.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inverse_elements.pdf
	open inverse_elements.pdf

o:
	make open

defs: inverse_elements.tex
	defs inverse_elements.tex

spell:
	aspell -c inverse_elements.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f inverse_elements.*
	make reset
	bbk_sheet -mode mk > Makefile