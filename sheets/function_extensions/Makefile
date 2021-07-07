# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make function_extensions.pdf

function_extensions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > function_extensions.tex

function_extensions.pdf: ../../*.tex ../../trademark.pdf *.tex function_extensions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode function_extensions.tex
	../../latexrun function_extensions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: function_extensions.pdf
	open function_extensions.pdf

o:
	make open

terms: function_extensions.tex
	bbk_sheet -mode ts function_extensions.tex

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
	rm -f function_extensions.*
	make reset
	bbk_sheet -mode mk > Makefile