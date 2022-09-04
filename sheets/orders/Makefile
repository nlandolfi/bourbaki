# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make orders.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

orders.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > orders.tex

orders.pdf: ../../*.tex ../../trademark.pdf *.tex orders.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode orders.tex
	../../latexrun orders.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: orders.pdf
	open orders.pdf

o:
	make open

terms: orders.tex
	bbk_sheet -mode ts orders.tex

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
	rm -f orders.*
	make reset
	bbk_sheet -mode mk > Makefile