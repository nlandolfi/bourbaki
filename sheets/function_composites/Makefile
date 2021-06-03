# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make function_composites.pdf

function_composites.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > function_composites.tex

function_composites.pdf: ../../*.tex ../../trademark.pdf *.tex function_composites.tex
	# pdflatex --file-line-error -interaction=nonstopmode function_composites.tex
	../../latexrun function_composites.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: function_composites.pdf
	open function_composites.pdf

o:
	make open

terms: function_composites.tex
	bbk_sheet -mode ts function_composites.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f function_composites.*
	make reset
	bbk_sheet -mode mk > Makefile