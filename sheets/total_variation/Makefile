# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make total_variation.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

total_variation.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > total_variation.tex

total_variation.pdf: ../../*.tex ../../trademark.pdf *.tex total_variation.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode total_variation.tex
	../../latexrun total_variation.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: total_variation.pdf
	open total_variation.pdf

o:
	make open

terms: total_variation.tex
	bbk sheet -mode ts total_variation.tex

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
	rm -f total_variation.*
	make reset
	bbk sheet -mode mk > Makefile