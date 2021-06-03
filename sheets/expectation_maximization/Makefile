# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make expectation_maximization.pdf

expectation_maximization.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > expectation_maximization.tex

expectation_maximization.pdf: ../../*.tex ../../trademark.pdf *.tex expectation_maximization.tex
	# pdflatex --file-line-error -interaction=nonstopmode expectation_maximization.tex
	../../latexrun expectation_maximization.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: expectation_maximization.pdf
	open expectation_maximization.pdf

o:
	make open

terms: expectation_maximization.tex
	bbk_sheet -mode ts expectation_maximization.tex

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
	rm -f expectation_maximization.*
	make reset
	bbk_sheet -mode mk > Makefile