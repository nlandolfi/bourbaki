# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make variance.pdf

variance.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > variance.tex

variance.pdf: ../../*.tex ../../trademark.pdf *.tex variance.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode variance.tex
	../../latexrun variance.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: variance.pdf
	open variance.pdf

o:
	make open

terms: variance.tex
	bbk_sheet -mode ts variance.tex

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
	rm -f variance.*
	make reset
	bbk_sheet -mode mk > Makefile