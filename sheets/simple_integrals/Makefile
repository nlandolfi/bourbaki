# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make simple_integrals.pdf

simple_integrals.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > simple_integrals.tex

simple_integrals.pdf: ../../*.tex ../../trademark.pdf *.tex simple_integrals.tex
	# pdflatex --file-line-error -interaction=nonstopmode simple_integrals.tex
	../../latexrun simple_integrals.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: simple_integrals.pdf
	open simple_integrals.pdf

o:
	make open

terms: simple_integrals.tex
	bbk_sheet -mode ts simple_integrals.tex

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
	rm -f simple_integrals.*
	make reset
	bbk_sheet -mode mk > Makefile