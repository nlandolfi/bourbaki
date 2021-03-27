# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make integrable_function_spaces.pdf
	make graph.pdf

integrable_function_spaces.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > integrable_function_spaces.tex

integrable_function_spaces.pdf: ../../*.tex ../../trademark.pdf *.tex integrable_function_spaces.tex
	pdflatex --file-line-error -interaction=nonstopmode integrable_function_spaces.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integrable_function_spaces.pdf
	open integrable_function_spaces.pdf

o:
	make open

defs: integrable_function_spaces.tex
	defs integrable_function_spaces.tex

spell:
	aspell -c integrable_function_spaces.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f integrable_function_spaces.*
	make reset
	bbk_sheet -mode mk > Makefile