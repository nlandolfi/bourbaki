# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make loss_functions.pdf
	make graph.pdf

loss_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > loss_functions.tex

loss_functions.pdf: ../../*.tex ../../trademark.pdf *.tex loss_functions.tex
	pdflatex --file-line-error -interaction=nonstopmode loss_functions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: loss_functions.pdf
	open loss_functions.pdf

o:
	make open

defs: loss_functions.tex
	defs loss_functions.tex

spell:
	aspell -c loss_functions.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f loss_functions.*
	make reset
	bbk_sheet -mode mk > Makefile