# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make sequences.pdf
	make graph.pdf

sequences.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > sequences.tex

sequences.pdf: ../../*.tex ../../trademark.pdf *.tex sequences.tex
	pdflatex --file-line-error -interaction=nonstopmode sequences.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: sequences.pdf
	open sequences.pdf

o:
	make open

defs: sequences.tex
	defs sequences.tex

spell:
	aspell -c sequences.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm sequences.*
	make reset
	bbk_sheet -mode mk > Makefile