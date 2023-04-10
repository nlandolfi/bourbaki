# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complete_inner_product_decomposition.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

complete_inner_product_decomposition.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > complete_inner_product_decomposition.tex

complete_inner_product_decomposition.pdf: ../../*.tex ../../trademark.pdf *.tex complete_inner_product_decomposition.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complete_inner_product_decomposition.tex
	../../latexrun complete_inner_product_decomposition.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complete_inner_product_decomposition.pdf
	open complete_inner_product_decomposition.pdf

o:
	make open

terms: complete_inner_product_decomposition.tex
	bbk sheet -mode ts complete_inner_product_decomposition.tex

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
	rm -f complete_inner_product_decomposition.*
	make reset
	bbk sheet -mode mk > Makefile