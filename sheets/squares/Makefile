# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make squares.pdf
	make graph.pdf

squares.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > squares.tex

squares.pdf: ../../*.tex ../../trademark.pdf *.tex squares.tex
	# pdflatex --file-line-error -interaction=nonstopmode squares.tex
	../../latexrun squares.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: squares.pdf
	open squares.pdf

o:
	make open

terms: squares.tex
	bbk_sheet -mode ts squares.tex

spell:
	aspell -c squares.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f squares.*
	make reset
	bbk_sheet -mode mk > Makefile