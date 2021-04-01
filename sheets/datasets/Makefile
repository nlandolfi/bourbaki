# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make datasets.pdf
	make graph.pdf

datasets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > datasets.tex

datasets.pdf: ../../*.tex ../../trademark.pdf *.tex datasets.tex
	pdflatex --file-line-error -interaction=nonstopmode datasets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: datasets.pdf
	open datasets.pdf

o:
	make open

terms: datasets.tex
	bbk_sheet -mode ts datasets.tex

spell:
	aspell -c datasets.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f datasets.*
	make reset
	bbk_sheet -mode mk > Makefile