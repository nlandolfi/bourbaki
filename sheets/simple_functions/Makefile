# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make simple_functions.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

simple_functions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > simple_functions.tex

simple_functions.pdf: ../../*.tex ../../trademark.pdf *.tex simple_functions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode simple_functions.tex
	../../latexrun simple_functions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: simple_functions.pdf
	open simple_functions.pdf

o:
	make open

terms: simple_functions.tex
	bbk_sheet -mode ts simple_functions.tex

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
	rm -f simple_functions.*
	make reset
	bbk_sheet -mode mk > Makefile