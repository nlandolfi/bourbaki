# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make partially_pivoted_row_reductions.pdf

partially_pivoted_row_reductions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > partially_pivoted_row_reductions.tex

partially_pivoted_row_reductions.pdf: ../../*.tex ../../trademark.pdf *.tex partially_pivoted_row_reductions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode partially_pivoted_row_reductions.tex
	../../latexrun partially_pivoted_row_reductions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: partially_pivoted_row_reductions.pdf
	open partially_pivoted_row_reductions.pdf

o:
	make open

terms: partially_pivoted_row_reductions.tex
	bbk_sheet -mode ts partially_pivoted_row_reductions.tex

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
	rm -f partially_pivoted_row_reductions.*
	make reset
	bbk_sheet -mode mk > Makefile