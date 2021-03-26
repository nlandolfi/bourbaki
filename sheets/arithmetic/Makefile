# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make arithmetic.pdf
	make graph.pdf

arithmetic.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > arithmetic.tex

arithmetic.pdf: ../../*.tex ../../trademark.pdf *.tex arithmetic.tex
	pdflatex --file-line-error -interaction=nonstopmode arithmetic.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: arithmetic.pdf
	open arithmetic.pdf

o:
	make open

defs: arithmetic.tex
	defs arithmetic.tex

spell:
	aspell -c arithmetic.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm arithmetic.*
	make reset
	bbk_sheet -mode mk > Makefile