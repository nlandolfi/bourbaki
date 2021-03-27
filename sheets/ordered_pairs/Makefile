# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make ordered_pairs.pdf
	make graph.pdf

ordered_pairs.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > ordered_pairs.tex

ordered_pairs.pdf: ../../*.tex ../../trademark.pdf *.tex ordered_pairs.tex
	pdflatex --file-line-error -interaction=nonstopmode ordered_pairs.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: ordered_pairs.pdf
	open ordered_pairs.pdf

o:
	make open

defs: ordered_pairs.tex
	defs ordered_pairs.tex

spell:
	aspell -c ordered_pairs.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f ordered_pairs.*
	make reset
	bbk_sheet -mode mk > Makefile