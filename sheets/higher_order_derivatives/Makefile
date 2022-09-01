# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make higher_order_derivatives.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

higher_order_derivatives.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > higher_order_derivatives.tex

higher_order_derivatives.pdf: ../../*.tex ../../trademark.pdf *.tex higher_order_derivatives.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode higher_order_derivatives.tex
	../../latexrun higher_order_derivatives.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: higher_order_derivatives.pdf
	open higher_order_derivatives.pdf

o:
	make open

terms: higher_order_derivatives.tex
	bbk_sheet -mode ts higher_order_derivatives.tex

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
	rm -f higher_order_derivatives.*
	make reset
	bbk_sheet -mode mk > Makefile