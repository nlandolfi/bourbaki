# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make partitions.pdf
	make graph.pdf

partitions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > partitions.tex

partitions.pdf: ../../*.tex ../../trademark.pdf *.tex partitions.tex
	pdflatex --file-line-error -interaction=nonstopmode partitions.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: partitions.pdf
	open partitions.pdf

o:
	make open

defs: partitions.tex
	defs partitions.tex

spell:
	aspell -c partitions.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm partitions.*
	make reset
	bbk_sheet -mode mk > Makefile