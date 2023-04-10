# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make simple_integral_additivity.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

simple_integral_additivity.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > simple_integral_additivity.tex

simple_integral_additivity.pdf: ../../*.tex ../../trademark.pdf *.tex simple_integral_additivity.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode simple_integral_additivity.tex
	../../latexrun simple_integral_additivity.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: simple_integral_additivity.pdf
	open simple_integral_additivity.pdf

o:
	make open

terms: simple_integral_additivity.tex
	bbk sheet -mode ts simple_integral_additivity.tex

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
	rm -f simple_integral_additivity.*
	make reset
	bbk sheet -mode mk > Makefile