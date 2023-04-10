# generated automatically by bbk sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make quasiconcave_functions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

quasiconcave_functions.tex: sheet.tex macros.tex
	bbk sheet -mode c -in sheet.tex > quasiconcave_functions.tex

quasiconcave_functions.pdf: ../../*.tex ../../trademark.pdf *.tex quasiconcave_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode quasiconcave_functions.tex
	../../latexrun quasiconcave_functions.tex
	make terms

graph.csv: sheet.tex
	bbk sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../bbk/static/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: quasiconcave_functions.pdf
	open quasiconcave_functions.pdf

o:
	make open

terms: quasiconcave_functions.tex
	bbk sheet -mode ts quasiconcave_functions.tex

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
	rm -f quasiconcave_functions.*
	make reset
	bbk sheet -mode mk > Makefile