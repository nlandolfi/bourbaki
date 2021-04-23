# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make total_orders.pdf
	make graph.pdf

total_orders.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > total_orders.tex

total_orders.pdf: ../../*.tex ../../trademark.pdf *.tex total_orders.tex
	# pdflatex --file-line-error -interaction=nonstopmode total_orders.tex
	../../latexrun total_orders.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: total_orders.pdf
	open total_orders.pdf

o:
	make open

terms: total_orders.tex
	bbk_sheet -mode ts total_orders.tex

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
	rm -f total_orders.*
	make reset
	bbk_sheet -mode mk > Makefile