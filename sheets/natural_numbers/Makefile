# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make natural_numbers.pdf
	make graph.pdf

natural_numbers.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_numbers.tex

natural_numbers.pdf: ../../*.tex ../../trademark.pdf *.tex natural_numbers.tex
	pdflatex --file-line-error -interaction=nonstopmode natural_numbers.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_numbers.pdf
	open natural_numbers.pdf

o:
	make open

defs: natural_numbers.tex
	defs natural_numbers.tex

spell:
	aspell -c natural_numbers.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f natural_numbers.*
	make reset
	bbk_sheet -mode mk > Makefile