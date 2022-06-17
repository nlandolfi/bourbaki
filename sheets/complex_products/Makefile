# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_products.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

complex_products.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_products.tex

complex_products.pdf: ../../*.tex ../../trademark.pdf *.tex complex_products.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_products.tex
	../../latexrun complex_products.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_products.pdf
	open complex_products.pdf

o:
	make open

terms: complex_products.tex
	bbk_sheet -mode ts complex_products.tex

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
	rm -f complex_products.*
	make reset
	bbk_sheet -mode mk > Makefile