# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make complex_integrals.pdf
	make graph.pdf

complex_integrals.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_integrals.tex

complex_integrals.pdf: ../../*.tex ../../trademark.pdf *.tex complex_integrals.tex
	pdflatex --file-line-error -interaction=nonstopmode complex_integrals.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_integrals.pdf
	open complex_integrals.pdf

o:
	make open

terms: complex_integrals.tex
	bbk_sheet -mode ts complex_integrals.tex

spell:
	aspell -c complex_integrals.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f complex_integrals.*
	make reset
	bbk_sheet -mode mk > Makefile