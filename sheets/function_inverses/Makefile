# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make function_inverses.pdf

function_inverses.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > function_inverses.tex

function_inverses.pdf: ../../*.tex ../../trademark.pdf *.tex function_inverses.tex
	# pdflatex --file-line-error -interaction=nonstopmode function_inverses.tex
	../../latexrun function_inverses.tex
	make terms

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

terms: function_inverses.tex
	bbk_sheet -mode ts function_inverses.tex

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
	rm -f function_inverses.*
	make reset
	bbk_sheet -mode mk > Makefile