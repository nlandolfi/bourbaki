# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make set_inclusion.pdf
	make graph.pdf

set_inclusion.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_inclusion.tex

set_inclusion.pdf: ../../*.tex ../../trademark.pdf *.tex set_inclusion.tex
	# pdflatex --file-line-error -interaction=nonstopmode set_inclusion.tex
	../../latexrun set_inclusion.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_inclusion.pdf
	open set_inclusion.pdf

o:
	make open

terms: set_inclusion.tex
	bbk_sheet -mode ts set_inclusion.tex

spell:
	aspell -c set_inclusion.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f set_inclusion.*
	make reset
	bbk_sheet -mode mk > Makefile