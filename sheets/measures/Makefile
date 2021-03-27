# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make measures.pdf
	make graph.pdf

measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > measures.tex

measures.pdf: ../../*.tex ../../trademark.pdf *.tex measures.tex
	pdflatex --file-line-error -interaction=nonstopmode measures.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: measures.pdf
	open measures.pdf

o:
	make open

defs: measures.tex
	defs measures.tex

spell:
	aspell -c measures.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f measures.*
	make reset
	bbk_sheet -mode mk > Makefile