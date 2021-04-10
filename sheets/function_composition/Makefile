# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make function_composition.pdf
	make graph.pdf

function_composition.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > function_composition.tex

function_composition.pdf: ../../*.tex ../../trademark.pdf *.tex function_composition.tex
	# pdflatex --file-line-error -interaction=nonstopmode function_composition.tex
	../../latexrun function_composition.tex
	make terms

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

terms: function_composition.tex
	bbk_sheet -mode ts function_composition.tex

spell:
	aspell -c function_composition.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f function_composition.*
	make reset
	bbk_sheet -mode mk > Makefile