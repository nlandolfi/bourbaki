# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make function_inverses.pdf
	make graph.pdf

function_inverses.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > function_inverses.tex

function_inverses.pdf: ../../*.tex ../../trademark.pdf *.tex function_inverses.tex
	pdflatex --file-line-error -interaction=nonstopmode function_inverses.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: function_inverses.pdf
	open function_inverses.pdf

o:
	make open

defs: function_inverses.tex
	defs function_inverses.tex

spell:
	aspell -c function_inverses.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm function_inverses.*
	make reset
	bbk_sheet -mode mk > Makefile