# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make supervised_learning.pdf
	make graph.pdf

supervised_learning.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > supervised_learning.tex

supervised_learning.pdf: ../../*.tex ../../trademark.pdf *.tex supervised_learning.tex
	pdflatex --file-line-error -interaction=nonstopmode supervised_learning.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: supervised_learning.pdf
	open supervised_learning.pdf

o:
	make open

terms: supervised_learning.tex
	bbk_sheet -mode ts supervised_learning.tex

spell:
	aspell -c supervised_learning.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f supervised_learning.*
	make reset
	bbk_sheet -mode mk > Makefile