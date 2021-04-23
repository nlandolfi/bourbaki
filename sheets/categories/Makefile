# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make categories.pdf
	make graph.pdf

categories.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > categories.tex

categories.pdf: ../../*.tex ../../trademark.pdf *.tex categories.tex
	# pdflatex --file-line-error -interaction=nonstopmode categories.tex
	../../latexrun categories.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: categories.pdf
	open categories.pdf

o:
	make open

terms: categories.tex
	bbk_sheet -mode ts categories.tex

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
	bbk_sheet -mode mk > Makefile