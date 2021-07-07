# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make set_specification.pdf

set_specification.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_specification.tex

set_specification.pdf: ../../*.tex ../../trademark.pdf *.tex set_specification.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode set_specification.tex
	../../latexrun set_specification.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_specification.pdf
	open set_specification.pdf

o:
	make open

terms: set_specification.tex
	bbk_sheet -mode ts set_specification.tex

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
	rm -f set_specification.*
	make reset
	bbk_sheet -mode mk > Makefile