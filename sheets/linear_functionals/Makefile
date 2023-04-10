# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make linear_functionals.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

linear_functionals.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > linear_functionals.tex

linear_functionals.pdf: ../../*.tex ../../trademark.pdf *.tex linear_functionals.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode linear_functionals.tex
	../../latexrun linear_functionals.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_functionals.pdf
	open linear_functionals.pdf

o:
	make open

terms: linear_functionals.tex
	bbk sheet -mode ts linear_functionals.tex

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
	rm -f linear_functionals.*
	make reset
	bbk sheet -mode mk > Makefile