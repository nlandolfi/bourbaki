# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make convex_sets.pdf
	make graph.pdf

convex_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > convex_sets.tex

convex_sets.pdf: ../../*.tex ../../trademark.pdf *.tex convex_sets.tex
	pdflatex --file-line-error -interaction=nonstopmode convex_sets.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: convex_sets.pdf
	open convex_sets.pdf

o:
	make open

defs: convex_sets.tex
	defs convex_sets.tex

spell:
	aspell -c convex_sets.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f convex_sets.*
	make reset
	bbk_sheet -mode mk > Makefile