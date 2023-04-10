# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make objects.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

objects.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > objects.tex

objects.pdf: ../../*.tex ../../trademark.pdf *.tex objects.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode objects.tex
	../../latexrun objects.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: objects.pdf
	open objects.pdf

o:
	make open

terms: objects.tex
	bbk sheet -mode ts objects.tex

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
	rm -f objects.*
	make reset
	bbk sheet -mode mk > Makefile