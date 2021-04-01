# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make direct_products.pdf
	make graph.pdf

direct_products.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > direct_products.tex

direct_products.pdf: ../../*.tex ../../trademark.pdf *.tex direct_products.tex
	pdflatex --file-line-error -interaction=nonstopmode direct_products.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: direct_products.pdf
	open direct_products.pdf

o:
	make open

terms: direct_products.tex
	bbk_sheet -mode ts direct_products.tex

spell:
	aspell -c direct_products.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f direct_products.*
	make reset
	bbk_sheet -mode mk > Makefile