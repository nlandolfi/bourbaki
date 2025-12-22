# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make monoid_products.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

monoid_products.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > monoid_products.tex

monoid_products.pdf: ../../*.tex ../../trademark.pdf *.tex monoid_products.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode monoid_products.tex
	../../latexrun monoid_products.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monoid_products.pdf
	open monoid_products.pdf

o:
	make open

terms: monoid_products.tex
	bbk sheet -mode ts monoid_products.tex

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
	rm -f monoid_products.*
	make reset
	bbk sheet -mode mk > Makefile