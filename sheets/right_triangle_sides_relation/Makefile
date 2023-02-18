# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make right_triangle_sides_relation.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

right_triangle_sides_relation.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > right_triangle_sides_relation.tex

right_triangle_sides_relation.pdf: ../../*.tex ../../trademark.pdf *.tex right_triangle_sides_relation.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode right_triangle_sides_relation.tex
	../../latexrun right_triangle_sides_relation.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: right_triangle_sides_relation.pdf
	open right_triangle_sides_relation.pdf

o:
	make open

terms: right_triangle_sides_relation.tex
	bbk_sheet -mode ts right_triangle_sides_relation.tex

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
	rm -f right_triangle_sides_relation.*
	make reset
	bbk_sheet -mode mk > Makefile