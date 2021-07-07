# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make set_intersections.pdf

set_intersections.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > set_intersections.tex

set_intersections.pdf: ../../*.tex ../../trademark.pdf *.tex set_intersections.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode set_intersections.tex
	../../latexrun set_intersections.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: set_intersections.pdf
	open set_intersections.pdf

o:
	make open

terms: set_intersections.tex
	bbk_sheet -mode ts set_intersections.tex

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
	rm -f set_intersections.*
	make reset
	bbk_sheet -mode mk > Makefile