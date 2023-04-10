# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_inner_products.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

complex_inner_products.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > complex_inner_products.tex

complex_inner_products.pdf: ../../*.tex ../../trademark.pdf *.tex complex_inner_products.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_inner_products.tex
	../../latexrun complex_inner_products.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_inner_products.pdf
	open complex_inner_products.pdf

o:
	make open

terms: complex_inner_products.tex
	bbk sheet -mode ts complex_inner_products.tex

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
	rm -f complex_inner_products.*
	make reset
	bbk sheet -mode mk > Makefile