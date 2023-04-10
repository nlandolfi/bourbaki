# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make characteristic_functions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

characteristic_functions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > characteristic_functions.tex

characteristic_functions.pdf: ../../*.tex ../../trademark.pdf *.tex characteristic_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode characteristic_functions.tex
	../../latexrun characteristic_functions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: characteristic_functions.pdf
	open characteristic_functions.pdf

o:
	make open

terms: characteristic_functions.tex
	bbk sheet -mode ts characteristic_functions.tex

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
	rm -f characteristic_functions.*
	make reset
	bbk sheet -mode mk > Makefile