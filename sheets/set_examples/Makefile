# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make set_examples.pdf

set_examples.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_examples.tex

set_examples.pdf: ../../*.tex ../../trademark.pdf *.tex set_examples.tex
	# pdflatex --file-line-error -interaction=nonstopmode set_examples.tex
	../../latexrun set_examples.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_examples.pdf
	open set_examples.pdf

o:
	make open

terms: set_examples.tex
	bbk_sheet -mode ts set_examples.tex

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
	rm -f set_examples.*
	make reset
	bbk_sheet -mode mk > Makefile