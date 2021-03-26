# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make closest_point_property.pdf
	make graph.pdf

closest_point_property.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > closest_point_property.tex

closest_point_property.pdf: ../../*.tex ../../trademark.pdf *.tex closest_point_property.tex
	pdflatex --file-line-error -interaction=nonstopmode closest_point_property.tex
	make defs

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: closest_point_property.pdf
	open closest_point_property.pdf

o:
	make open

defs: closest_point_property.tex
	defs closest_point_property.tex

spell:
	aspell -c closest_point_property.tex

s:
	make spell

reset:
	rm graph.pdf
	rm graph.csv
	rm graph.graphviz

clean:
	make reset

remake:
	rm closest_point_property.*
	make reset
	bbk_sheet -mode mk > Makefile