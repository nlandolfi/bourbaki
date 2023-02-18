# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make space_vector_angles.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

space_vector_angles.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > space_vector_angles.tex

space_vector_angles.pdf: ../../*.tex ../../trademark.pdf *.tex space_vector_angles.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode space_vector_angles.tex
	../../latexrun space_vector_angles.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: space_vector_angles.pdf
	open space_vector_angles.pdf

o:
	make open

terms: space_vector_angles.tex
	bbk_sheet -mode ts space_vector_angles.tex

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
	rm -f space_vector_angles.*
	make reset
	bbk_sheet -mode mk > Makefile