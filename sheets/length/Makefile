# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make length.pdf
	make graph.pdf

length.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > length.tex

length.pdf: ../../*.tex ../../trademark.pdf *.tex length.tex
	pdflatex --file-line-error -interaction=nonstopmode length.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: length.pdf
	open length.pdf

o:
	make open

defs: length.tex
	defs length.tex

spell:
	aspell -c length.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm length.*
	make reset
	bbk_sheet -mode mk > Makefile