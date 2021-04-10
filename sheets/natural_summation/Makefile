# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make natural_summation.pdf
	make graph.pdf

natural_summation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_summation.tex

natural_summation.pdf: ../../*.tex ../../trademark.pdf *.tex natural_summation.tex
	# pdflatex --file-line-error -interaction=nonstopmode natural_summation.tex
	../../latexrun natural_summation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_summation.pdf
	open natural_summation.pdf

o:
	make open

terms: natural_summation.tex
	bbk_sheet -mode ts natural_summation.tex

spell:
	aspell -c natural_summation.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f natural_summation.*
	make reset
	bbk_sheet -mode mk > Makefile