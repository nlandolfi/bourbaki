# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make function_composition.pdf
	make graph.pdf

function_composition.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > function_composition.tex

function_composition.pdf: ../../*.tex ../../trademark.pdf *.tex function_composition.tex
	pdflatex --file-line-error -interaction=nonstopmode function_composition.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: function_composition.pdf
	open function_composition.pdf

o:
	make open

defs: function_composition.tex
	defs function_composition.tex

spell:
	aspell -c function_composition.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm function_composition.*
	make reset
	bbk_sheet -mode mk > Makefile