# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make family_products_and_unions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

family_products_and_unions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > family_products_and_unions.tex

family_products_and_unions.pdf: ../../*.tex ../../trademark.pdf *.tex family_products_and_unions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode family_products_and_unions.tex
	../../latexrun family_products_and_unions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: family_products_and_unions.pdf
	open family_products_and_unions.pdf

o:
	make open

terms: family_products_and_unions.tex
	bbk sheet -mode ts family_products_and_unions.tex

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
	rm -f family_products_and_unions.*
	make reset
	bbk sheet -mode mk > Makefile