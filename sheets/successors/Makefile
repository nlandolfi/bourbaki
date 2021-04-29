# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make successors.pdf
	make graph.pdf

successors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > successors.tex

successors.pdf: ../../*.tex ../../trademark.pdf *.tex successors.tex
	# pdflatex --file-line-error -interaction=nonstopmode successors.tex
	../../latexrun successors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: successors.pdf
	open successors.pdf

o:
	make open

terms: successors.tex
	bbk_sheet -mode ts successors.tex

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
	rm -f successors.*
	make reset
	bbk_sheet -mode mk > Makefile