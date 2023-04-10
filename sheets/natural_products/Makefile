# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_products.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

natural_products.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > natural_products.tex

natural_products.pdf: ../../*.tex ../../trademark.pdf *.tex natural_products.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_products.tex
	../../latexrun natural_products.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_products.pdf
	open natural_products.pdf

o:
	make open

terms: natural_products.tex
	bbk_sheet -mode ts natural_products.tex

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
	rm -f natural_products.*
	make reset
	bbk_sheet -mode mk > Makefile