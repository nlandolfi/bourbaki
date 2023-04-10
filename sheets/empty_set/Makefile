# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make empty_set.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

empty_set.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > empty_set.tex

empty_set.pdf: ../../*.tex ../../trademark.pdf *.tex empty_set.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode empty_set.tex
	../../latexrun empty_set.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: empty_set.pdf
	open empty_set.pdf

o:
	make open

terms: empty_set.tex
	bbk sheet -mode ts empty_set.tex

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
	rm -f empty_set.*
	make reset
	bbk sheet -mode mk > Makefile