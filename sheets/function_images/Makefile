# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make function_images.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

function_images.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > function_images.tex

function_images.pdf: ../../*.tex ../../trademark.pdf *.tex function_images.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode function_images.tex
	../../latexrun function_images.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: function_images.pdf
	open function_images.pdf

o:
	make open

terms: function_images.tex
	bbk sheet -mode ts function_images.tex

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
	rm -f function_images.*
	make reset
	bbk sheet -mode mk > Makefile