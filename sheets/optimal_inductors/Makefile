# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make optimal_inductors.pdf

optimal_inductors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > optimal_inductors.tex

optimal_inductors.pdf: ../../*.tex ../../trademark.pdf *.tex optimal_inductors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode optimal_inductors.tex
	../../latexrun optimal_inductors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: optimal_inductors.pdf
	open optimal_inductors.pdf

o:
	make open

terms: optimal_inductors.tex
	bbk_sheet -mode ts optimal_inductors.tex

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
	rm -f optimal_inductors.*
	make reset
	bbk_sheet -mode mk > Makefile