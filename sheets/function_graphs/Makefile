# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make function_graphs.pdf
	make graph.pdf

function_graphs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > function_graphs.tex

function_graphs.pdf: ../../*.tex ../../trademark.pdf *.tex function_graphs.tex
	# pdflatex --file-line-error -interaction=nonstopmode function_graphs.tex
	../../latexrun function_graphs.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: function_graphs.pdf
	open function_graphs.pdf

o:
	make open

terms: function_graphs.tex
	bbk_sheet -mode ts function_graphs.tex

spell:
	aspell -c function_graphs.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f function_graphs.*
	make reset
	bbk_sheet -mode mk > Makefile