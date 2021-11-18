# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make polar_coordinates.pdf

polar_coordinates.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > polar_coordinates.tex

polar_coordinates.pdf: ../../*.tex ../../trademark.pdf *.tex polar_coordinates.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode polar_coordinates.tex
	../../latexrun polar_coordinates.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: polar_coordinates.pdf
	open polar_coordinates.pdf

o:
	make open

terms: polar_coordinates.tex
	bbk_sheet -mode ts polar_coordinates.tex

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
	rm -f polar_coordinates.*
	make reset
	bbk_sheet -mode mk > Makefile