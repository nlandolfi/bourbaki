# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make natural_addition.pdf
	make graph.pdf

natural_addition.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_addition.tex

natural_addition.pdf: ../../*.tex ../../trademark.pdf *.tex natural_addition.tex
	# pdflatex --file-line-error -interaction=nonstopmode natural_addition.tex
	../../latexrun natural_addition.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_addition.pdf
	open natural_addition.pdf

o:
	make open

terms: natural_addition.tex
	bbk_sheet -mode ts natural_addition.tex

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
	rm -f natural_addition.*
	make reset
	bbk_sheet -mode mk > Makefile