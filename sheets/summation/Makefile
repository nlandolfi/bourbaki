# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make summation.pdf
	make graph.pdf

summation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > summation.tex

summation.pdf: ../../*.tex ../../trademark.pdf *.tex summation.tex
	pdflatex --file-line-error -interaction=nonstopmode summation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: summation.pdf
	open summation.pdf

o:
	make open

terms: summation.tex
	bbk_sheet -mode ts summation.tex

spell:
	aspell -c summation.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f summation.*
	make reset
	bbk_sheet -mode mk > Makefile