# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make inner_product_linear_functional_representations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

inner_product_linear_functional_representations.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > inner_product_linear_functional_representations.tex

inner_product_linear_functional_representations.pdf: ../../*.tex ../../trademark.pdf *.tex inner_product_linear_functional_representations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode inner_product_linear_functional_representations.tex
	../../latexrun inner_product_linear_functional_representations.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inner_product_linear_functional_representations.pdf
	open inner_product_linear_functional_representations.pdf

o:
	make open

terms: inner_product_linear_functional_representations.tex
	bbk_sheet -mode ts inner_product_linear_functional_representations.tex

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
	rm -f inner_product_linear_functional_representations.*
	make reset
	bbk_sheet -mode mk > Makefile