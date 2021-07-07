# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make partial_derivatives.pdf

partial_derivatives.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > partial_derivatives.tex

partial_derivatives.pdf: ../../*.tex ../../trademark.pdf *.tex partial_derivatives.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode partial_derivatives.tex
	../../latexrun partial_derivatives.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: partial_derivatives.pdf
	open partial_derivatives.pdf

o:
	make open

terms: partial_derivatives.tex
	bbk_sheet -mode ts partial_derivatives.tex

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
	rm -f partial_derivatives.*
	make reset
	bbk_sheet -mode mk > Makefile