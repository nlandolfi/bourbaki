# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make function_properties.pdf
	make graph.pdf

function_properties.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > function_properties.tex

function_properties.pdf: ../../*.tex ../../trademark.pdf *.tex function_properties.tex
	pdflatex --file-line-error -interaction=nonstopmode function_properties.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: function_properties.pdf
	open function_properties.pdf

o:
	make open

defs: function_properties.tex
	defs function_properties.tex

spell:
	aspell -c function_properties.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f function_properties.*
	make reset
	bbk_sheet -mode mk > Makefile