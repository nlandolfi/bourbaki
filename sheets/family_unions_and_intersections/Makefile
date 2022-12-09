# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make family_unions_and_intersections.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

family_unions_and_intersections.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > family_unions_and_intersections.tex

family_unions_and_intersections.pdf: ../../*.tex ../../trademark.pdf *.tex family_unions_and_intersections.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode family_unions_and_intersections.tex
	../../latexrun family_unions_and_intersections.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: family_unions_and_intersections.pdf
	open family_unions_and_intersections.pdf

o:
	make open

terms: family_unions_and_intersections.tex
	bbk_sheet -mode ts family_unions_and_intersections.tex

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
	rm -f family_unions_and_intersections.*
	make reset
	bbk_sheet -mode mk > Makefile