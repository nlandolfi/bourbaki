# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make transpose.pdf
	make graph.pdf

transpose.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > transpose.tex

transpose.pdf: ../../*.tex ../../trademark.pdf *.tex transpose.tex
	pdflatex --file-line-error -interaction=nonstopmode transpose.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: transpose.pdf
	open transpose.pdf

o:
	make open

defs: transpose.tex
	defs transpose.tex

spell:
	aspell -c transpose.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm transpose.*
	make reset
	bbk_sheet -mode mk > Makefile