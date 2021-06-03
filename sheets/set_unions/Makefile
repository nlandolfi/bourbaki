# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make set_unions.pdf

set_unions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_unions.tex

set_unions.pdf: ../../*.tex ../../trademark.pdf *.tex set_unions.tex
	# pdflatex --file-line-error -interaction=nonstopmode set_unions.tex
	../../latexrun set_unions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_unions.pdf
	open set_unions.pdf

o:
	make open

terms: set_unions.tex
	bbk_sheet -mode ts set_unions.tex

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
	rm -f set_unions.*
	make reset
	bbk_sheet -mode mk > Makefile