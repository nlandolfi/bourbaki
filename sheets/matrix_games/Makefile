# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make matrix_games.pdf
	make graph.pdf

matrix_games.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > matrix_games.tex

matrix_games.pdf: ../../*.tex ../../trademark.pdf *.tex matrix_games.tex
	# pdflatex --file-line-error -interaction=nonstopmode matrix_games.tex
	../../latexrun matrix_games.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: matrix_games.pdf
	open matrix_games.pdf

o:
	make open

terms: matrix_games.tex
	bbk_sheet -mode ts matrix_games.tex

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
	rm -f matrix_games.*
	make reset
	bbk_sheet -mode mk > Makefile