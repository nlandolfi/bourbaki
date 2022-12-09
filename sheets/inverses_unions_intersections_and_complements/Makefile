# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make inverses_unions_intersections_and_complements.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

inverses_unions_intersections_and_complements.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > inverses_unions_intersections_and_complements.tex

inverses_unions_intersections_and_complements.pdf: ../../*.tex ../../trademark.pdf *.tex inverses_unions_intersections_and_complements.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode inverses_unions_intersections_and_complements.tex
	../../latexrun inverses_unions_intersections_and_complements.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: inverses_unions_intersections_and_complements.pdf
	open inverses_unions_intersections_and_complements.pdf

o:
	make open

terms: inverses_unions_intersections_and_complements.tex
	bbk_sheet -mode ts inverses_unions_intersections_and_complements.tex

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
	rm -f inverses_unions_intersections_and_complements.*
	make reset
	bbk_sheet -mode mk > Makefile