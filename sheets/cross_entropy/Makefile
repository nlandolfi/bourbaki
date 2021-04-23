# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make cross_entropy.pdf
	make graph.pdf

cross_entropy.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cross_entropy.tex

cross_entropy.pdf: ../../*.tex ../../trademark.pdf *.tex cross_entropy.tex
	# pdflatex --file-line-error -interaction=nonstopmode cross_entropy.tex
	../../latexrun cross_entropy.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cross_entropy.pdf
	open cross_entropy.pdf

o:
	make open

terms: cross_entropy.tex
	bbk_sheet -mode ts cross_entropy.tex

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
	rm -f cross_entropy.*
	make reset
	bbk_sheet -mode mk > Makefile