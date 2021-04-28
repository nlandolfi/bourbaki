# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make natural_difference.pdf
	make graph.pdf

natural_difference.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > natural_difference.tex

natural_difference.pdf: ../../*.tex ../../trademark.pdf *.tex natural_difference.tex
	# pdflatex --file-line-error -interaction=nonstopmode natural_difference.tex
	../../latexrun natural_difference.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_difference.pdf
	open natural_difference.pdf

o:
	make open

terms: natural_difference.tex
	bbk_sheet -mode ts natural_difference.tex

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
	rm -f natural_difference.*
	make reset
	bbk_sheet -mode mk > Makefile