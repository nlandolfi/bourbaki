# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make image_measures.pdf
	make graph.pdf

image_measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > image_measures.tex

image_measures.pdf: ../../*.tex ../../trademark.pdf *.tex image_measures.tex
	pdflatex --file-line-error -interaction=nonstopmode image_measures.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: image_measures.pdf
	open image_measures.pdf

o:
	make open

defs: image_measures.tex
	defs image_measures.tex

spell:
	aspell -c image_measures.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f image_measures.*
	make reset
	bbk_sheet -mode mk > Makefile