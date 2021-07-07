# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make expectation.pdf

expectation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > expectation.tex

expectation.pdf: ../../*.tex ../../trademark.pdf *.tex expectation.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode expectation.tex
	../../latexrun expectation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: expectation.pdf
	open expectation.pdf

o:
	make open

terms: expectation.tex
	bbk_sheet -mode ts expectation.tex

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
	rm -f expectation.*
	make reset
	bbk_sheet -mode mk > Makefile