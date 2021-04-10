# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make density_selection.pdf
	make graph.pdf

density_selection.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > density_selection.tex

density_selection.pdf: ../../*.tex ../../trademark.pdf *.tex density_selection.tex
	# pdflatex --file-line-error -interaction=nonstopmode density_selection.tex
	../../latexrun density_selection.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: density_selection.pdf
	open density_selection.pdf

o:
	make open

terms: density_selection.tex
	bbk_sheet -mode ts density_selection.tex

spell:
	aspell -c density_selection.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f density_selection.*
	make reset
	bbk_sheet -mode mk > Makefile