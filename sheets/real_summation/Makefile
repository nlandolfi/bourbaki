# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_summation.pdf

real_summation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_summation.tex

real_summation.pdf: ../../*.tex ../../trademark.pdf *.tex real_summation.tex
	# pdflatex --file-line-error -interaction=nonstopmode real_summation.tex
	../../latexrun real_summation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_summation.pdf
	open real_summation.pdf

o:
	make open

terms: real_summation.tex
	bbk_sheet -mode ts real_summation.tex

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
	rm -f real_summation.*
	make reset
	bbk_sheet -mode mk > Makefile