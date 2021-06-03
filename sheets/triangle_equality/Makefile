# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make triangle_equality.pdf

triangle_equality.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > triangle_equality.tex

triangle_equality.pdf: ../../*.tex ../../trademark.pdf *.tex triangle_equality.tex
	# pdflatex --file-line-error -interaction=nonstopmode triangle_equality.tex
	../../latexrun triangle_equality.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: triangle_equality.pdf
	open triangle_equality.pdf

o:
	make open

terms: triangle_equality.tex
	bbk_sheet -mode ts triangle_equality.tex

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
	rm -f triangle_equality.*
	make reset
	bbk_sheet -mode mk > Makefile