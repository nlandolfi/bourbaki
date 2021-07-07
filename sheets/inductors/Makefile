# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make inductors.pdf

inductors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > inductors.tex

inductors.pdf: ../../*.tex ../../trademark.pdf *.tex inductors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode inductors.tex
	../../latexrun inductors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inductors.pdf
	open inductors.pdf

o:
	make open

terms: inductors.tex
	bbk_sheet -mode ts inductors.tex

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
	rm -f inductors.*
	make reset
	bbk_sheet -mode mk > Makefile