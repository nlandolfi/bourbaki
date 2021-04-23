# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make inner_product_norms.pdf
	make graph.pdf

inner_product_norms.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > inner_product_norms.tex

inner_product_norms.pdf: ../../*.tex ../../trademark.pdf *.tex inner_product_norms.tex
	# pdflatex --file-line-error -interaction=nonstopmode inner_product_norms.tex
	../../latexrun inner_product_norms.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inner_product_norms.pdf
	open inner_product_norms.pdf

o:
	make open

terms: inner_product_norms.tex
	bbk_sheet -mode ts inner_product_norms.tex

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
	rm -f inner_product_norms.*
	make reset
	bbk_sheet -mode mk > Makefile