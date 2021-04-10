# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make function_images.pdf
	make graph.pdf

function_images.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > function_images.tex

function_images.pdf: ../../*.tex ../../trademark.pdf *.tex function_images.tex
	# pdflatex --file-line-error -interaction=nonstopmode function_images.tex
	../../latexrun function_images.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: function_images.pdf
	open function_images.pdf

o:
	make open

terms: function_images.tex
	bbk_sheet -mode ts function_images.tex

spell:
	aspell -c function_images.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f function_images.*
	make reset
	bbk_sheet -mode mk > Makefile