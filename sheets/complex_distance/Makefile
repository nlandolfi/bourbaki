# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make complex_distance.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

complex_distance.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > complex_distance.tex

complex_distance.pdf: ../../*.tex ../../trademark.pdf *.tex complex_distance.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode complex_distance.tex
	../../latexrun complex_distance.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: complex_distance.pdf
	open complex_distance.pdf

o:
	make open

terms: complex_distance.tex
	bbk sheet -mode ts complex_distance.tex

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
	rm -f complex_distance.*
	make reset
	bbk sheet -mode mk > Makefile