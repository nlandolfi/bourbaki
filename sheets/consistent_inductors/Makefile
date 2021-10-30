# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make consistent_inductors.pdf

consistent_inductors.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > consistent_inductors.tex

consistent_inductors.pdf: ../../*.tex ../../trademark.pdf *.tex consistent_inductors.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode consistent_inductors.tex
	../../latexrun consistent_inductors.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: consistent_inductors.pdf
	open consistent_inductors.pdf

o:
	make open

terms: consistent_inductors.tex
	bbk_sheet -mode ts consistent_inductors.tex

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
	rm -f consistent_inductors.*
	make reset
	bbk_sheet -mode mk > Makefile