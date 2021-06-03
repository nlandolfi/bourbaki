# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make distribution_selection.pdf

distribution_selection.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > distribution_selection.tex

distribution_selection.pdf: ../../*.tex ../../trademark.pdf *.tex distribution_selection.tex
	# pdflatex --file-line-error -interaction=nonstopmode distribution_selection.tex
	../../latexrun distribution_selection.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: distribution_selection.pdf
	open distribution_selection.pdf

o:
	make open

terms: distribution_selection.tex
	bbk_sheet -mode ts distribution_selection.tex

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
	rm -f distribution_selection.*
	make reset
	bbk_sheet -mode mk > Makefile