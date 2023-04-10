# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make natural_order.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

natural_order.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > natural_order.tex

natural_order.pdf: ../../*.tex ../../trademark.pdf *.tex natural_order.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode natural_order.tex
	../../latexrun natural_order.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: natural_order.pdf
	open natural_order.pdf

o:
	make open

terms: natural_order.tex
	bbk_sheet -mode ts natural_order.tex

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
	rm -f natural_order.*
	make reset
	bbk_sheet -mode mk > Makefile