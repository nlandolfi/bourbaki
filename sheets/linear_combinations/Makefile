# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make linear_combinations.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

linear_combinations.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > linear_combinations.tex

linear_combinations.pdf: ../../*.tex ../../trademark.pdf *.tex linear_combinations.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode linear_combinations.tex
	../../latexrun linear_combinations.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_combinations.pdf
	open linear_combinations.pdf

o:
	make open

terms: linear_combinations.tex
	bbk sheet -mode ts linear_combinations.tex

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
	rm -f linear_combinations.*
	make reset
	bbk sheet -mode mk > Makefile