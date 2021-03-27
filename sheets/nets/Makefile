# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make nets.pdf
	make graph.pdf

nets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > nets.tex

nets.pdf: ../../*.tex ../../trademark.pdf *.tex nets.tex
	pdflatex --file-line-error -interaction=nonstopmode nets.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: nets.pdf
	open nets.pdf

o:
	make open

defs: nets.tex
	defs nets.tex

spell:
	aspell -c nets.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f nets.*
	make reset
	bbk_sheet -mode mk > Makefile