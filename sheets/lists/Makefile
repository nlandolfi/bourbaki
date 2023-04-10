# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make lists.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

lists.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > lists.tex

lists.pdf: ../../*.tex ../../trademark.pdf *.tex lists.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode lists.tex
	../../latexrun lists.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: lists.pdf
	open lists.pdf

o:
	make open

terms: lists.tex
	bbk sheet -mode ts lists.tex

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
	rm -f lists.*
	make reset
	bbk sheet -mode mk > Makefile