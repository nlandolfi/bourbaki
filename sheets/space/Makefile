# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make space.pdf
	make graph.pdf

space.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > space.tex

space.pdf: ../../*.tex ../../trademark.pdf *.tex space.tex
	# pdflatex --file-line-error -interaction=nonstopmode space.tex
	../../latexrun space.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: space.pdf
	open space.pdf

o:
	make open

terms: space.tex
	bbk_sheet -mode ts space.tex

spell:
	aspell -c space.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f space.*
	make reset
	bbk_sheet -mode mk > Makefile