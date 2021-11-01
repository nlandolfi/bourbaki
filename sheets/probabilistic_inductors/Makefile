# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make probabilistic_inductors.pdf

probabilistic_inductors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > probabilistic_inductors.tex

probabilistic_inductors.pdf: ../../*.tex ../../trademark.pdf *.tex probabilistic_inductors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode probabilistic_inductors.tex
	../../latexrun probabilistic_inductors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: probabilistic_inductors.pdf
	open probabilistic_inductors.pdf

o:
	make open

terms: probabilistic_inductors.tex
	bbk_sheet -mode ts probabilistic_inductors.tex

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
	rm -f probabilistic_inductors.*
	make reset
	bbk_sheet -mode mk > Makefile