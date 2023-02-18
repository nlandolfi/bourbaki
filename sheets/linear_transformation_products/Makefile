# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make linear_transformation_products.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

linear_transformation_products.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > linear_transformation_products.tex

linear_transformation_products.pdf: ../../*.tex ../../trademark.pdf *.tex linear_transformation_products.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode linear_transformation_products.tex
	../../latexrun linear_transformation_products.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_transformation_products.pdf
	open linear_transformation_products.pdf

o:
	make open

terms: linear_transformation_products.tex
	bbk_sheet -mode ts linear_transformation_products.tex

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
	rm -f linear_transformation_products.*
	make reset
	bbk_sheet -mode mk > Makefile