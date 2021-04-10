# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make set_equality.pdf
	make graph.pdf

set_equality.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_equality.tex

set_equality.pdf: ../../*.tex ../../trademark.pdf *.tex set_equality.tex
	# pdflatex --file-line-error -interaction=nonstopmode set_equality.tex
	../../latexrun set_equality.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_equality.pdf
	open set_equality.pdf

o:
	make open

terms: set_equality.tex
	bbk_sheet -mode ts set_equality.tex

spell:
	aspell -c set_equality.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f set_equality.*
	make reset
	bbk_sheet -mode mk > Makefile