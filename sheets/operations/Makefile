# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make operations.pdf
	make graph.pdf

operations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > operations.tex

operations.pdf: ../../*.tex ../../trademark.pdf *.tex operations.tex
	pdflatex --file-line-error -interaction=nonstopmode operations.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: operations.pdf
	open operations.pdf

o:
	make open

defs: operations.tex
	defs operations.tex

spell:
	aspell -c operations.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm operations.*
	make reset
	bbk_sheet -mode mk > Makefile