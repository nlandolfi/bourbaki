# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make dynamic_games.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

dynamic_games.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > dynamic_games.tex

dynamic_games.pdf: ../../*.tex ../../trademark.pdf *.tex dynamic_games.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode dynamic_games.tex
	../../latexrun dynamic_games.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: dynamic_games.pdf
	open dynamic_games.pdf

o:
	make open

terms: dynamic_games.tex
	bbk_sheet -mode ts dynamic_games.tex

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
	rm -f dynamic_games.*
	make reset
	bbk_sheet -mode mk > Makefile