# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make barycentric_coordinates.pdf

barycentric_coordinates.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > barycentric_coordinates.tex

barycentric_coordinates.pdf: ../../*.tex ../../trademark.pdf *.tex barycentric_coordinates.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode barycentric_coordinates.tex
	../../latexrun barycentric_coordinates.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: barycentric_coordinates.pdf
	open barycentric_coordinates.pdf

o:
	make open

terms: barycentric_coordinates.tex
	bbk_sheet -mode ts barycentric_coordinates.tex

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
	rm -f barycentric_coordinates.*
	make reset
	bbk_sheet -mode mk > Makefile