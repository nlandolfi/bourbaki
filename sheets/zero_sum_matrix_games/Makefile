# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make zero_sum_matrix_games.pdf
	make graph.pdf

zero_sum_matrix_games.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > zero_sum_matrix_games.tex

zero_sum_matrix_games.pdf: ../../*.tex ../../trademark.pdf *.tex zero_sum_matrix_games.tex
	pdflatex --file-line-error -interaction=nonstopmode zero_sum_matrix_games.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: zero_sum_matrix_games.pdf
	open zero_sum_matrix_games.pdf

o:
	make open

defs: zero_sum_matrix_games.tex
	defs zero_sum_matrix_games.tex

spell:
	aspell -c zero_sum_matrix_games.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f zero_sum_matrix_games.*
	make reset
	bbk_sheet -mode mk > Makefile