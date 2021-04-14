# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make approximation.pdf
	make graph.pdf

approximation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > approximation.tex

approximation.pdf: ../../*.tex ../../trademark.pdf *.tex approximation.tex
	# pdflatex --file-line-error -interaction=nonstopmode approximation.tex
	../../latexrun approximation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: approximation.pdf
	open approximation.pdf

o:
	make open

terms: approximation.tex
	bbk_sheet -mode ts approximation.tex

spell:
	aspell -c approximation.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f approximation.*
	make reset
	bbk_sheet -mode mk > Makefile