# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make introduction.pdf
	make graph.pdf

introduction.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > introduction.tex

introduction.pdf: ../../*.tex ../../trademark.pdf *.tex introduction.tex
	pdflatex --file-line-error -interaction=nonstopmode introduction.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: introduction.pdf
	open introduction.pdf

o:
	make open

defs: introduction.tex
	defs introduction.tex

spell:
	aspell -c introduction.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm introduction.*
	make reset
	bbk_sheet -mode mk > Makefile