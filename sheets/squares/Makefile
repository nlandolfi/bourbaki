# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make squares.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

squares.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > squares.tex

squares.pdf: ../../*.tex ../../trademark.pdf *.tex squares.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode squares.tex
	../../latexrun squares.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: squares.pdf
	open squares.pdf

o:
	make open

terms: squares.tex
	bbk sheet -mode ts squares.tex

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
	rm -f squares.*
	make reset
	bbk sheet -mode mk > Makefile