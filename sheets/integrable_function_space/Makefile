# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make integrable_function_space.pdf
	make graph.pdf

integrable_function_space.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > integrable_function_space.tex

integrable_function_space.pdf: ../../*.tex ../../trademark.pdf *.tex integrable_function_space.tex
	# pdflatex --file-line-error -interaction=nonstopmode integrable_function_space.tex
	../../latexrun integrable_function_space.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: integrable_function_space.pdf
	open integrable_function_space.pdf

o:
	make open

terms: integrable_function_space.tex
	bbk_sheet -mode ts integrable_function_space.tex

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
	rm -f integrable_function_space.*
	make reset
	bbk_sheet -mode mk > Makefile