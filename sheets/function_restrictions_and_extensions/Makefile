# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make function_restrictions_and_extensions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

function_restrictions_and_extensions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > function_restrictions_and_extensions.tex

function_restrictions_and_extensions.pdf: ../../*.tex ../../trademark.pdf *.tex function_restrictions_and_extensions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode function_restrictions_and_extensions.tex
	../../latexrun function_restrictions_and_extensions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: function_restrictions_and_extensions.pdf
	open function_restrictions_and_extensions.pdf

o:
	make open

terms: function_restrictions_and_extensions.tex
	bbk sheet -mode ts function_restrictions_and_extensions.tex

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
	rm -f function_restrictions_and_extensions.*
	make reset
	bbk sheet -mode mk > Makefile