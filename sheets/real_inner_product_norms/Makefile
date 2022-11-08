# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_inner_product_norms.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

real_inner_product_norms.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > real_inner_product_norms.tex

real_inner_product_norms.pdf: ../../*.tex ../../trademark.pdf *.tex real_inner_product_norms.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_inner_product_norms.tex
	../../latexrun real_inner_product_norms.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_inner_product_norms.pdf
	open real_inner_product_norms.pdf

o:
	make open

terms: real_inner_product_norms.tex
	bbk_sheet -mode ts real_inner_product_norms.tex

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
	rm -f real_inner_product_norms.*
	make reset
	bbk_sheet -mode mk > Makefile