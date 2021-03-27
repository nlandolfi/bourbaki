# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make two-player_one-shot_games.pdf
	make graph.pdf

two-player_one-shot_games.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > two-player_one-shot_games.tex

two-player_one-shot_games.pdf: ../../*.tex ../../trademark.pdf *.tex two-player_one-shot_games.tex
	pdflatex --file-line-error -interaction=nonstopmode two-player_one-shot_games.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: two-player_one-shot_games.pdf
	open two-player_one-shot_games.pdf

o:
	make open

defs: two-player_one-shot_games.tex
	defs two-player_one-shot_games.tex

spell:
	aspell -c two-player_one-shot_games.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f two-player_one-shot_games.*
	make reset
	bbk_sheet -mode mk > Makefile