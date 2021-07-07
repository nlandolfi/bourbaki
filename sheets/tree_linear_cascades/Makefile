# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make tree_linear_cascades.pdf

tree_linear_cascades.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > tree_linear_cascades.tex

tree_linear_cascades.pdf: ../../*.tex ../../trademark.pdf *.tex tree_linear_cascades.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode tree_linear_cascades.tex
	../../latexrun tree_linear_cascades.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: tree_linear_cascades.pdf
	open tree_linear_cascades.pdf

o:
	make open

terms: tree_linear_cascades.tex
	bbk_sheet -mode ts tree_linear_cascades.tex

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
	rm -f tree_linear_cascades.*
	make reset
	bbk_sheet -mode mk > Makefile