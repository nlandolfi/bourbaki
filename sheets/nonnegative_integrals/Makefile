# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make nonnegative_integrals.pdf
	make graph.pdf

nonnegative_integrals.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > nonnegative_integrals.tex

nonnegative_integrals.pdf: ../../*.tex ../../trademark.pdf *.tex nonnegative_integrals.tex
	pdflatex --file-line-error -interaction=nonstopmode nonnegative_integrals.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: nonnegative_integrals.pdf
	open nonnegative_integrals.pdf

o:
	make open

defs: nonnegative_integrals.tex
	defs nonnegative_integrals.tex

spell:
	aspell -c nonnegative_integrals.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f nonnegative_integrals.*
	make reset
	bbk_sheet -mode mk > Makefile