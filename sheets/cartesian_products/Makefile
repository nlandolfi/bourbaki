# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make cartesian_products.pdf

cartesian_products.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cartesian_products.tex

cartesian_products.pdf: ../../*.tex ../../trademark.pdf *.tex cartesian_products.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode cartesian_products.tex
	../../latexrun cartesian_products.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cartesian_products.pdf
	open cartesian_products.pdf

o:
	make open

terms: cartesian_products.tex
	bbk_sheet -mode ts cartesian_products.tex

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
	rm -f cartesian_products.*
	make reset
	bbk_sheet -mode mk > Makefile