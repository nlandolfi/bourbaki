# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make probability.pdf
	make graph.pdf

probability.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probability.tex

probability.pdf: ../../*.tex ../../trademark.pdf *.tex probability.tex
	pdflatex --file-line-error -interaction=nonstopmode probability.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probability.pdf
	open probability.pdf

o:
	make open

defs: probability.tex
	defs probability.tex

spell:
	aspell -c probability.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f probability.*
	make reset
	bbk_sheet -mode mk > Makefile