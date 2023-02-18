# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_circular_coordinates.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

complex_circular_coordinates.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > complex_circular_coordinates.tex

complex_circular_coordinates.pdf: ../../*.tex ../../trademark.pdf *.tex complex_circular_coordinates.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_circular_coordinates.tex
	../../latexrun complex_circular_coordinates.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_circular_coordinates.pdf
	open complex_circular_coordinates.pdf

o:
	make open

terms: complex_circular_coordinates.tex
	bbk_sheet -mode ts complex_circular_coordinates.tex

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
	rm -f complex_circular_coordinates.*
	make reset
	bbk_sheet -mode mk > Makefile