# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make fair_coins.pdf

fair_coins.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > fair_coins.tex

fair_coins.pdf: ../../*.tex ../../trademark.pdf *.tex fair_coins.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode fair_coins.tex
	../../latexrun fair_coins.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: fair_coins.pdf
	open fair_coins.pdf

o:
	make open

terms: fair_coins.tex
	bbk_sheet -mode ts fair_coins.tex

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
	rm -f fair_coins.*
	make reset
	bbk_sheet -mode mk > Makefile