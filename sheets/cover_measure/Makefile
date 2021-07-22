# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make cover_measure.pdf

cover_measure.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > cover_measure.tex

cover_measure.pdf: ../../*.tex ../../trademark.pdf *.tex cover_measure.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode cover_measure.tex
	../../latexrun cover_measure.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: cover_measure.pdf
	open cover_measure.pdf

o:
	make open

terms: cover_measure.tex
	bbk_sheet -mode ts cover_measure.tex

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
	rm -f cover_measure.*
	make reset
	bbk_sheet -mode mk > Makefile