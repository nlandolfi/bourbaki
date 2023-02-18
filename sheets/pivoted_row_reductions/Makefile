# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make pivoted_row_reductions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

pivoted_row_reductions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > pivoted_row_reductions.tex

pivoted_row_reductions.pdf: ../../*.tex ../../trademark.pdf *.tex pivoted_row_reductions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode pivoted_row_reductions.tex
	../../latexrun pivoted_row_reductions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: pivoted_row_reductions.pdf
	open pivoted_row_reductions.pdf

o:
	make open

terms: pivoted_row_reductions.tex
	bbk_sheet -mode ts pivoted_row_reductions.tex

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
	rm -f pivoted_row_reductions.*
	make reset
	bbk_sheet -mode mk > Makefile