# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make discrete_sets.pdf

discrete_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > discrete_sets.tex

discrete_sets.pdf: ../../*.tex ../../trademark.pdf *.tex discrete_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode discrete_sets.tex
	../../latexrun discrete_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: discrete_sets.pdf
	open discrete_sets.pdf

o:
	make open

terms: discrete_sets.tex
	bbk_sheet -mode ts discrete_sets.tex

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
	rm -f discrete_sets.*
	make reset
	bbk_sheet -mode mk > Makefile