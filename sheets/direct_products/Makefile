# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make direct_products.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

direct_products.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > direct_products.tex

direct_products.pdf: ../../*.tex ../../trademark.pdf *.tex direct_products.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode direct_products.tex
	../../latexrun direct_products.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: direct_products.pdf
	open direct_products.pdf

o:
	make open

terms: direct_products.tex
	bbk sheet -mode ts direct_products.tex

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
	rm -f direct_products.*
	make reset
	bbk sheet -mode mk > Makefile