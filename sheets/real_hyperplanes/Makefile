# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make real_hyperplanes.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

real_hyperplanes.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > real_hyperplanes.tex

real_hyperplanes.pdf: ../../*.tex ../../trademark.pdf *.tex real_hyperplanes.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode real_hyperplanes.tex
	../../latexrun real_hyperplanes.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: real_hyperplanes.pdf
	open real_hyperplanes.pdf

o:
	make open

terms: real_hyperplanes.tex
	bbk sheet -mode ts real_hyperplanes.tex

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
	rm -f real_hyperplanes.*
	make reset
	bbk sheet -mode mk > Makefile