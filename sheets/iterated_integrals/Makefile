# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make iterated_integrals.pdf

iterated_integrals.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > iterated_integrals.tex

iterated_integrals.pdf: ../../*.tex ../../trademark.pdf *.tex iterated_integrals.tex
	# pdflatex --file-line-error -interaction=nonstopmode iterated_integrals.tex
	../../latexrun iterated_integrals.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: iterated_integrals.pdf
	open iterated_integrals.pdf

o:
	make open

terms: iterated_integrals.tex
	bbk_sheet -mode ts iterated_integrals.tex

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
	rm -f iterated_integrals.*
	make reset
	bbk_sheet -mode mk > Makefile