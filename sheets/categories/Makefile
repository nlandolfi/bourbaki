# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make categories.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

categories.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > categories.tex

categories.pdf: ../../*.tex ../../trademark.pdf *.tex categories.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode categories.tex
	../../latexrun categories.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: categories.pdf
	open categories.pdf

o:
	make open

terms: categories.tex
	bbk sheet -mode ts categories.tex

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
	rm -f categories.*
	make reset
	bbk sheet -mode mk > Makefile