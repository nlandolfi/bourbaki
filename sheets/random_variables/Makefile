# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make random_variables.pdf
	make graph.pdf

random_variables.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > random_variables.tex

random_variables.pdf: ../../*.tex ../../trademark.pdf *.tex random_variables.tex
	pdflatex --file-line-error -interaction=nonstopmode random_variables.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: random_variables.pdf
	open random_variables.pdf

o:
	make open

terms: random_variables.tex
	bbk_sheet -mode ts random_variables.tex

spell:
	aspell -c random_variables.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f random_variables.*
	make reset
	bbk_sheet -mode mk > Makefile