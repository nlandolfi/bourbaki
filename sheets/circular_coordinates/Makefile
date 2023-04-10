# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make circular_coordinates.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

circular_coordinates.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > circular_coordinates.tex

circular_coordinates.pdf: ../../*.tex ../../trademark.pdf *.tex circular_coordinates.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode circular_coordinates.tex
	../../latexrun circular_coordinates.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: circular_coordinates.pdf
	open circular_coordinates.pdf

o:
	make open

terms: circular_coordinates.tex
	bbk sheet -mode ts circular_coordinates.tex

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
	rm -f circular_coordinates.*
	make reset
	bbk sheet -mode mk > Makefile