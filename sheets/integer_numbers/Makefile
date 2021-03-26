# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make integer_numbers.pdf
	make graph.pdf

integer_numbers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > integer_numbers.tex

integer_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex integer_numbers.tex
	pdflatex --file-line-error -interaction=nonstopmode integer_numbers.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integer_numbers.pdf
	open integer_numbers.pdf

o:
	make open

defs: integer_numbers.tex
	defs integer_numbers.tex

spell:
	aspell -c integer_numbers.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm integer_numbers.*
	make reset
	bbk_sheet -mode mk > Makefile