# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make space_inner_product.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

space_inner_product.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > space_inner_product.tex

space_inner_product.pdf: ../../*.tex ../../trademark.pdf *.tex space_inner_product.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode space_inner_product.tex
	../../latexrun space_inner_product.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: space_inner_product.pdf
	open space_inner_product.pdf

o:
	make open

terms: space_inner_product.tex
	bbk_sheet -mode ts space_inner_product.tex

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
	rm -f space_inner_product.*
	make reset
	bbk_sheet -mode mk > Makefile