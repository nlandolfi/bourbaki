# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make geometry.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

geometry.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > geometry.tex

geometry.pdf: ../../*.tex ../../trademark.pdf *.tex geometry.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode geometry.tex
	../../latexrun geometry.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: geometry.pdf
	open geometry.pdf

o:
	make open

terms: geometry.tex
	bbk sheet -mode ts geometry.tex

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
	rm -f geometry.*
	make reset
	bbk sheet -mode mk > Makefile