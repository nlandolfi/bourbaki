# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make input_designs.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

input_designs.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > input_designs.tex

input_designs.pdf: ../../*.tex ../../trademark.pdf *.tex input_designs.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode input_designs.tex
	../../latexrun input_designs.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: input_designs.pdf
	open input_designs.pdf

o:
	make open

terms: input_designs.tex
	bbk sheet -mode ts input_designs.tex

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
	rm -f input_designs.*
	make reset
	bbk sheet -mode mk > Makefile