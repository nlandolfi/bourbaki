# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make natural_multiplication.pdf
	make graph.pdf

natural_multiplication.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_multiplication.tex

natural_multiplication.pdf: ../../*.tex ../../trademark.pdf *.tex natural_multiplication.tex
	# pdflatex --file-line-error -interaction=nonstopmode natural_multiplication.tex
	../../latexrun natural_multiplication.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_multiplication.pdf
	open natural_multiplication.pdf

o:
	make open

terms: natural_multiplication.tex
	bbk_sheet -mode ts natural_multiplication.tex

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
	rm -f natural_multiplication.*
	make reset
	bbk_sheet -mode mk > Makefile