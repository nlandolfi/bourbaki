# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make combined_orders.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

combined_orders.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > combined_orders.tex

combined_orders.pdf: ../../*.tex ../../trademark.pdf *.tex combined_orders.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode combined_orders.tex
	../../latexrun combined_orders.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: combined_orders.pdf
	open combined_orders.pdf

o:
	make open

terms: combined_orders.tex
	bbk sheet -mode ts combined_orders.tex

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
	rm -f combined_orders.*
	make reset
	bbk sheet -mode mk > Makefile